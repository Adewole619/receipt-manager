package database

import (
	"database/sql"
	"fmt"

	_ "modernc.org/sqlite"
)

var DB *sql.DB

func Connect() error {

	db, err := sql.Open("sqlite", "receipt.db")
	if err != nil {
		return err
	}

	DB = db

	fmt.Println("Connected to SQLite successfully!")

	return nil
}
