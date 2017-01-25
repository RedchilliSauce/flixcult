package data

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
)

const exists = `SELECT COUNT (*) FROM sqlite_master 
        WHERE type = 'table' AND name = user;`

var queries = [...]string{

	`CREATE TABLE user(
		user_name     VARCHAR( 255 ) NOT NULL,
		first_name    VARCHAR( 255 ),
		second_name   VARCHAR( 255 ),
		email         VARCHAR( 255 ) NOT NULL,
		PRIMARY KEY( user_name ),
		UNIQUE(email)
    );`,

	`CREATE TABLE user_password(
    	user_name   VARCHAR( 255 ) NOT NULL,
    	hash        VARCHAR( 255 ) NOT NULL,
    	PRIMARY KEY( user_name ),
    	FOREIGN KEY( user_name ) REFERENCES user( user_name ) 
			ON DELETE CASCADE,
    );`,

	`CREATE TABLE group(
		group_id	VARCHAR( 256 ) NOT NULL
    	name        VARCHAR( 256 ) NOT NULL,
    	owner       VARCHAR( 256 ) NOT NULL,
    	description TEXT NOT NULL,
        visiblity   VARCHAR( 256 ) NOT NULL
    	PRIMARY KEY( group_id ),
    	FOREIGN KEY( owner ) REFERENCES user( user_name )
    );`,

	`CREATE TABLE user_to_group(
    	group_id    VARCHAR( 256 ) NOT NULL,
    	user_name   VARCHAR( 256 ) NOT NULL,
    	FOREIGN KEY( group_id ) REFERENCES user_group( group_id )
			ON DELETE CASCADE,
    	FOREIGN KEY( user_name ) REFERENCES user( user_name ),
    	PRIMARY KEY( group_id, user_name )
			ON DELETE CASCADE
    );`,
}

//PostgresStore - datasource implemented using postgress
type PostgresStore struct {
	*sqlx.DB
}

//Options - options for connecting to sqlite database
type Options struct {
	UserName string
	Password string
	DBName   string
}

//Init - initializes the datastore
func Init(options *Options) (*PostgresStore, error) {
	dbArgs := fmt.Sprintf(
		"user=%s password=%s dbname=%s sslmode=disable",
		options.UserName,
		options.Password,
		options.DBName)
	db, err := sqlx.Connect("postgres", dbArgs)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	defer db.Close()
	var sqliteDB *PostgresStore
	if err = db.Ping(); err == nil {
		row := db.QueryRow(exists)
		var count int
		err = row.Scan(&count)
		if err == nil {
			if count == 0 {
				sqliteDB, err = create(options, db)
			} else {
				sqliteDB, err = connect(options)
			}
		} else {
			log.Print(err)
			return nil, err
		}
	}
	return sqliteDB, err
}

//connect - connects to a sqlite database file
func connect(options *Options) (*PostgresStore, error) {
	dbArgs := fmt.Sprintf(
		"user=%s password=%s dbname=%s sslmode=disable",
		options.UserName,
		options.Password,
		options.DBName)
	mdb, err := sqlx.Connect("postgres", dbArgs)
	if err != nil {
		log.Print(err)
	} else if err = mdb.Ping(); err != nil {
		log.Printf("Could not connect to mysql database: %s", err)
	} else {
		log.Print("Database opened successfuly")
	}
	return &PostgresStore{mdb}, err
}

//create - connects to a sqlite database file and creates schema
func create(options *Options, db *sqlx.DB) (*PostgresStore, error) {
	mdb, err := connect(options)
	if err == nil {
		for index, query := range queries {
			_, err = mdb.Exec(query)
			if err != nil {
				log.Printf(`Failed to create database, query %d failed: %s`,
					index, err)
				break
			}
		}
	}
	return mdb, err
}
