package db

import (
	"database/sql"
	"log"
	"time"
)

var DB *sql.DB

func InitDb() {
	DB, err := sql.Open("sqlite3", "api.db")
	if err != nil {
		log.Fatalln(err)
	}
	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)
	DB.SetConnMaxLifetime(5 * time.Minute)
	err = DB.Ping()
	if err != nil {
		log.Fatal("Database Not Reachable")
	}
}
