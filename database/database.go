package database

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func GetDatabase() *sql.DB {
	db_url := "gofan.db"
	db, err := sql.Open("sqlite3", db_url)
	if err != nil {
		fmt.Println(err)
	}
	return db
}
