package database

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "reports.db")
	if err != nil {
		log.Fatal(err)
	}

	createTable()
}

func createTable() {
	query := `CREATE TABLE IF NOT EXISTS reports (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		content TEXT NOT NULL,
		timestamp DATETIME DEFAULT CURRENT_TIMESTAMP
	)`
	_, err := DB.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
}
