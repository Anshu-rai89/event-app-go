package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var DB sql.DB

func InitDB() {
	DB, err := sql.Open("sqlite3", "api.db")

	if err != nil {
		panic("Couldn't connect to DB")
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)
}

func createTable() {
	createTableEvent := `
		CREATE TABLE IF NOT EXISTS events(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		description TEXT NOT NULL,
		location TEXT NOT NULL,
		dateTime DATETIME,
		user_id INTEGER
	)`

	_, err := DB.Exec(createTableEvent)

	if err != nil {
		panic("Couldn't create Table")
	}
}
