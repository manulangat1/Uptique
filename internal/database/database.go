package database

import (
	"database/sql"
	"os"

	_ "modernc.org/sqlite"
)

func Open() (*sql.DB, error) {
	if err := os.MkdirAll("./data", 0755); err != nil {
		return nil, err
	}

	db, err := sql.Open("sqlite", "./data/scheduler.db")
	if err != nil {
		return nil, err
	}

	if _, err := db.Exec("PRAGMA foreign_keys = ON"); err != nil {
		return nil, err
	}

	return db, nil
}
