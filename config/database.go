package config

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB
func ConnectDB(){
	db, err := sql.Open("mysql", "root:root@/gobimbingan")
	if err != nil {
		panic(err)
	}

	log.Println("Connected to database")
	DB = db
}