package database

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/go-sql-driver/mysql"
)

var db *sql.DB

func SqlConnect() (db *sql.DB) {
	if db == nil {
		cfg := mysql.Config{
			User:   "root",
			Passwd: "root@123",
			Net:    "tcp",
			Addr:   "127.0.0.1:3306",
			DBName: "RetailSale",
		}
		// Get a database handle.
		var err error
		db, err = sql.Open("mysql", cfg.FormatDSN())
		if err != nil {
			log.Printf("Error in opening mysql %v", err.Error())
			panic(err)
		}

		pingErr := db.Ping()
		if pingErr != nil {
			log.Printf("Error in creating connection %v", pingErr.Error())
			panic(pingErr)
		}
		fmt.Println("Connected to Mysql!")
	}
	return db

}
