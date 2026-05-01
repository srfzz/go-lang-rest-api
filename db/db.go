package db

import (
	"database/sql"
	"log"
	"log/slog"
	"time"

	_ "modernc.org/sqlite"
)

var DB *sql.DB

func InitDb() {
	var err error
	DB, err = sql.Open("sqlite", "api.db")
	if err != nil {
		log.Fatalln(err)
	}
	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)
	DB.SetConnMaxLifetime(5 * time.Minute)
	err = DB.Ping()
	if err != nil {
		log.Fatalf("Database Not Reachable %v", err.Error())
	}
	createTables()
}

func createTables() {
	createEventtabe := `
	CREATE table  if not exists events(
	id Integer primary key Autoincrement,
	title text not null,
	description text not null,
	location text not null,
	dateTime datetime not null,
	user_id Integer 
	)
	`
	_, err := DB.Exec(createEventtabe)
	if err != nil {
		log.Fatalf("Table not Created %v", err.Error())
		return
	}
	slog.Info("table created")
}
