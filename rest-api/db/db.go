package db

import (
	"database/sql"
	// _ means we need the import and must be part of final project, stops it being auto removed. It will be used by standard lib sql import
	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDb() {

	var err error
	//driver name sqlite3 means go-sqlite3 import is used
	//api.db - it will be created if it doesn't exist
	//open just opens db, nothing else - just allows connections to start
	DB, err = sql.Open("sqlite3", "api.db")

	if err != nil {
		panic("failed to connect to database")
	}

	//create a pool of connections
	DB.SetMaxOpenConns(10)
	//make 5 always idle and available, we can use up to 10
	DB.SetMaxIdleConns(5)

	createTables()
}

func createTables() {

	createUsersTable := `
	CREATE TABLE IF NOT EXISTS users (
	    id INTEGER PRIMARY KEY AUTOINCREMENT,
	    email TEXT NOT NULL UNIQUE,
	    password TEXT NOT NULL
	)
	`

	_, err := DB.Exec(createUsersTable)
	if err != nil {
		panic("could not cease create users table")
	}

	createEventsTable := `
	CREATE TABLE IF NOT EXISTS events (
    	id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		description TEXT NOT NULL,
		location TEXT NOT NULL,
		dateTime DATETIME NOT NULL,
		user_id INTEGER,
	    FOREIGN KEY(user_id) REFERENCES users(id)
	)`

	_, err = DB.Exec(createEventsTable)
	if err != nil {
		panic("could not create events table")
	}

	createRegistrationsTable := `
	CREATE TABLE IF NOT EXISTS registrations (
	    id INTEGER PRIMARY KEY AUTOINCREMENT,
	    event_id INTEGER,
	    user_id INTEGER,
	    FOREIGN KEY(event_id) REFERENCES events(id),
	    FOREIGN KEY(user_id) REFERENCES users(id)
	)
	`

	_, err = DB.Exec(createRegistrationsTable)

	if err != nil {
		panic("could not create registrations table, reason: " + err.Error())
	}
}
