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
		panic("Could not connect to database.")
	}

	DB.SetMaxOpenConns(10) // how many connections can be opened simultaneously
	DB.SetMaxIdleConns(5) // how many connections we want to keep open if no one is using these connections at the moment

	/* createTables() */
}
