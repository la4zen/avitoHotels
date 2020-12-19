package util

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func GetConnect() *sql.DB {
	db, err := sql.Open("sqlite3", "./database.db")
	if err != nil {
		fmt.Errorf("sql error")
	}
	return db
}
