package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/mattn/go-sqlite3"
	. "github.com/roobert/golang-error"
)

var DB *sql.DB

func OpenDB(db string) {
	var err error
	DB, err = sql.Open("sqlite3", db)
	CheckErr(err)
}

func DeleteDB(db string) {
	err := os.Remove(db)
	CheckErr(err)
}

func CreateTable(table, fields string) {
	query := fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s (%s)", table, fields)
	_, err := DB.Exec(query)
	CheckErr(err)
}
