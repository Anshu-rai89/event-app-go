package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "api.db")

	if err != nil {
		panic("Couldn't connect to DB")
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)
	createTable()
}

func createTable() {

	createUserTable := `
	CREATE TABLE IF NOT EXISTS users(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		email TEXT NOT NULL UNIQUE,
		name TEXT NOT NULL,
		password TEXT NOT NULL
	)`

	_, err := DB.Exec(createUserTable)

	createTableEvent := `
		CREATE TABLE IF NOT EXISTS events(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		description TEXT NOT NULL,
		location TEXT NOT NULL,
		dateTime DATETIME,
		user_id INTEGER,
		FOREIGN KEY(user_id) REFERENCES users(id)
	)`

	_, err = DB.Exec(createTableEvent)

	if err != nil {
		panic("Couldn't create Table")
	}

	createEventRegistrationsTable := `
	CREATE TABLE  IF NOT EXIST eventRegistrations(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		userId INTEGER NOT NULL,
		eventId INTEGER NOT NULL,
		FOREIGN KEY(userId) REFERENCES users(id),
		FOREIGN KEY(eventId) REFERENCES events(id)
	)`

	_, err = DB.Exec(createEventRegistrationsTable)

	if err != nil {
		panic("Couldn't create Table")
	}
}
