package mysql

import (
	_ "github.com/go-sql-driver/mysql" // MySQL driver

	"database/sql"
	"log"
)

// DB mysql global
var DB *sql.DB

// Connect to the MySQL database and return the db connection
func Connect() {
	var err error
	DB, err = sql.Open("mysql", "username:password@tcp(192.168.178.5:3306)/db")
	if err != nil {
		log.Fatal(err)
	}
}

// Disconnect to the MySQL database and return the db connection
func Disconnect(DB *sql.DB) {
	defer DB.Close()
}
