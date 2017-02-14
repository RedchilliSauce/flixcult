package data

import (
	"fmt"
	"log"

	"database/sql"

	"github.com/jmoiron/sqlx"
)

const exists = `SELECT EXISTS (
	SELECT table_name FROM information_schema.tables WHERE 
	table_schema='public' AND table_name='%s');`

//CreateQuery - Represents a query that creates a table
type CreateQuery struct {
	TableName   string
	QueryString string
}

var queries = [...]CreateQuery{
	CreateQuery{
		TableName: "fc_user",
		QueryString: `CREATE TABLE fc_user(
			user_name     VARCHAR( 255 ) NOT NULL,
			first_name    VARCHAR( 255 ),
			second_name   VARCHAR( 255 ),
			email         VARCHAR( 255 ) NOT NULL,
			PRIMARY KEY( user_name ),
			UNIQUE(email)
    	);`,
	},
	CreateQuery{
		TableName: "fc_user_password",
		QueryString: `CREATE TABLE fc_user_password(
    		user_name   VARCHAR( 255 ) NOT NULL,
    		hash        VARCHAR( 255 ) NOT NULL,
    		PRIMARY KEY( user_name ),
    		FOREIGN KEY( user_name ) REFERENCES fc_user( user_name ) 
				ON DELETE CASCADE
    	);`,
	},
	CreateQuery{
		TableName: "fc_group",
		QueryString: `CREATE TABLE fc_group(
			group_id	VARCHAR( 256 ) NOT NULL,
    		name        VARCHAR( 256 ) NOT NULL,
    		owner       VARCHAR( 256 ) NOT NULL,
    		description TEXT NOT NULL,
    	    visibility   VARCHAR( 256 ) NOT NULL,
    		PRIMARY KEY( group_id ),
    		FOREIGN KEY( owner ) REFERENCES fc_user( user_name )
    	);`,
	},
	CreateQuery{
		TableName: "fc_user_to_group",
		QueryString: `CREATE TABLE fc_user_to_group(
    		group_id    VARCHAR( 256 ) NOT NULL,
    		user_name   VARCHAR( 256 ) NOT NULL,
    		PRIMARY KEY( group_id, user_name ),
    		FOREIGN KEY( group_id ) REFERENCES fc_group( group_id ) 
				ON DELETE CASCADE,
    		FOREIGN KEY( user_name ) REFERENCES fc_user( user_name ) 
				ON DELETE CASCADE
    	);`,
	},
}

//PostgresStore - datasource implemented using postgress
type PostgresStore struct {
	*sqlx.DB
}

//PsqlOptions - options for connecting to sqlite database
type PsqlOptions struct {
	UserName string
	Password string
	Host     string
	Port     int
	DBName   string
}

//PsqlInit - initializes the datastore
func PsqlInit(options *PsqlOptions) (pgdb *PostgresStore, err error) {
	//postgres://<user>:<password>@<host>:<port>/<dbname>?<params>
	connInfo := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		options.UserName,
		options.Password,
		options.Host,
		options.Port,
		options.DBName)
	var db *sqlx.DB
	db, err = sqlx.Connect("postgres", connInfo)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	defer db.Close()

	if err = db.Ping(); err == nil {
		pgdb, err = create(options, db)
	}
	if err != nil {
		log.Println(err)
	}
	return pgdb, err
}

//connect - connects to a sqlite database file
func connect(options *PsqlOptions) (*PostgresStore, error) {
	connInfo := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		options.UserName,
		options.Password,
		options.Host,
		options.Port,
		options.DBName)
	mdb, err := sqlx.Connect("postgres", connInfo)
	if err != nil {
		log.Print(err)
	} else if err = mdb.Ping(); err != nil {
		log.Printf("Could not connect to postgress database: %s", err)
	} else {
		log.Print("Database opened successfuly")
	}
	return &PostgresStore{mdb}, err
}

//create - connects to a sqlite database file and creates schema
func create(options *PsqlOptions, db *sqlx.DB) (*PostgresStore, error) {
	mdb, err := connect(options)
	if err == nil {
		for _, query := range queries {
			if !tableExists(db, query.TableName) {
				_, err = mdb.Exec(query.QueryString)
				if err != nil {
					log.Printf(`Could not create table %s: %s`,
						query.TableName, err)
					// break
				} else {
					log.Printf("Created table %s", query.TableName)
				}
			} else {
				log.Printf("Table %s already exists, nothing to do!",
					query.TableName)
			}
		}
	}
	return mdb, err
}

func tableExists(db *sqlx.DB, tableName string) (has bool) {
	query := fmt.Sprintf(exists, tableName)
	rows := db.QueryRowx(query)
	err := rows.Err()
	if err == nil {
		err = rows.Scan(&has)
	}
	if err == sql.ErrNoRows {
		has = false
	} else if err != nil {
		log.Fatalf("Could not initialize database: %s", err)
		has = false
	}
	return has
}
