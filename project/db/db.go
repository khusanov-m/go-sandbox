package db

import (
	"database/sql"

	_ "github.com/ncruces/go-sqlite3/driver"
	_ "github.com/ncruces/go-sqlite3/embed" // Import sqlite3 driver, but don't use it directly
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "api.db") // Should end with .db

	if err != nil {
		panic("Could not connect to database!")
	}

	DB.SetMaxOpenConns(10)   // Set max number of connections open simultaneously 10
	DB.SetConnMaxIdleTime(5) // At least 5 connections should be open at all times - to be ready to process requests

	createTables()
}

func createTables() {
	createEventsTable := `
    CREATE TABLE IF NOT EXISTS events (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        name TEXT NOT NULL,
        description TEXT NOT NULL, 
        location TEXT NOT NULL,
        dateTime DATETIME NOT NULL,
        userId INTEGER
    )
    `

	_, err := DB.Exec(createEventsTable)
	if err != nil {
		panic("Could not create events table.")
	}
}
