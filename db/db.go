package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var DBClient *sql.DB

func InitDB(){
	var err error
	DBClient, err =sql.Open("sqlite3", "api.db")

	if err!=nil{
		panic("Could not connect to the database")
	}

	DBClient.SetMaxOpenConns(10);
	DBClient.SetConnMaxIdleTime(5);

	createTable()
}

func createTable(){

	createEventsTable := 
	`CREATE TABLE IF NOT EXISTS events(
	 	id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		description TEXT NOT NULL,
		location TEXT NOT NULL,
		dateTime INTEGER NOT NULL,
		userId INTEGER NOT NULL
	 )`

	_, err:= DBClient.Exec(createEventsTable)

	if err!=nil{
		panic("Could not create events table")
	}
}
