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

	createTables()
}

// Create the events table
func createEventTable(){
	createEventsTable := 
	`CREATE TABLE IF NOT EXISTS events(
	 	id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		description TEXT NOT NULL,
		location TEXT NOT NULL,
		dateTime INTEGER NOT NULL,
		userId INTEGER NOT NULL,
		FOREIGN KEY (userId) REFERENCES users(id)
	 )`

	_, err:= DBClient.Exec(createEventsTable)

	if err!=nil{
		panic("Could not create events table")
	}
}

// Create the users table
func createUserTable(){
	createUsersTable := 
	`CREATE TABLE IF NOT EXISTS users(
	 	id INTEGER PRIMARY KEY AUTOINCREMENT,
		email TEXT NOT NULL UNIQUE,
		password TEXT NOT NULL
	 )`

	_, err:= DBClient.Exec(createUsersTable)

	if err!=nil{
		panic("Could not create users table")
	}
}

// Registration table
func createRegistrationTable(){
	createRegistrationTable := 
	`CREATE TABLE IF NOT EXISTS registrations(
	 	id INTEGER PRIMARY KEY AUTOINCREMENT,
		userId INTEGER NOT NULL,
		eventId INTEGER NOT NULL,
		FOREIGN KEY (userId) REFERENCES users(id),
		FOREIGN KEY (eventId) REFERENCES events(id)
	 )`

	_, err:= DBClient.Exec(createRegistrationTable)

	if err!=nil{
		panic("Could not create registrations table")
	}
}



func createTables() {
	createUserTable()
	createEventTable()
	createRegistrationTable()
}
