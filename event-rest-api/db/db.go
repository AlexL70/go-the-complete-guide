package db

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "./events-api.db")
	if err != nil {
		panic(err)
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	createTables()
}

func createTables() {
	createUsersStmt := `
CREATE TABLE IF NOT EXISTS users(
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	email TEXT NOT NULL UNIQUE,
	password TEXT NOT NULL
);
`

	_, err := DB.Exec((createUsersStmt))
	if err != nil {
		panic(fmt.Errorf("Cannot create \"users\" table: %w", err))
	}

	createEventsStmt := `
CREATE TABLE IF NOT EXISTS events (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    description TEXT NOT NULL,
    location TEXT NOT NULL,
    date_time DATETIME NOT NULL,
    user_id INTEGER NOT NULL,
	FOREIGN KEY(user_id) REFERENCES users(id)
);
`
	_, err = DB.Exec(createEventsStmt)
	if err != nil {
		panic(fmt.Errorf("Cannot create \"events\" table: %w", err))
	}

	createRegistrationsStmt := `
CREATE TABLE IF NOT EXISTS registrations (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	event_id INTEGER NOT NULL,
	user_id INTEGER NOT NULL,
	FOREIGN KEY(event_id) REFERENCES events(id),
	FOREIGN KEY(user_id) REFERENCES users(id)
	UNIQUE(event_id, user_id)
);
`
	_, err = DB.Exec(createRegistrationsStmt)
	if err != nil {
		panic(fmt.Errorf("Cannot create \"registrations\" table: %w", err))
	}
}
