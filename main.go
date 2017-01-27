package main

import (
	"flag"
	"fmt"
	"log"
	"syscall"

	_ "github.com/lib/pq"
	"golang.org/x/crypto/ssh/terminal"

	"github.com/RedchilliSauce/flixcult/data"
)

func main() {

	var userName string
	var host string
	var port int
	var dbName string
	var password string

	flag.StringVar(&userName, "user-name", "", "Database user name")
	flag.StringVar(&host, "host", "localhost", "Database host")
	flag.StringVar(&dbName, "db-name", "flixcult", "Name of the database")
	flag.IntVar(&port, "port", 5432, "Database port")
	flag.StringVar(&password, "password", "",
		"Password, will be asked later if not provided")
	flag.Parse()
	if len(password) == 0 && len(userName) != 0 {
		fmt.Printf("Password for %s: ", userName)
		pbyte, err := terminal.ReadPassword(int(syscall.Stdin))
		if err != nil {
			fmt.Print(err)
		} else {
			password = string(pbyte)
		}
	} else if len(userName) == 0 {
		log.Fatal("Insufficient parameters for FlixCult to run!")
	}
	options := &data.PsqlOptions{
		UserName: userName,
		Password: password,
		Host:     host,
		Port:     port,
		DBName:   dbName}
	// fmt.Println(options)
	db, err := data.PsqlInit(options)
	if err == nil {
		data.SetStore(db)
	} else {
		log.Fatal("Failed to initialize database")
	}
}
