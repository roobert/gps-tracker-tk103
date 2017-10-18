package ui

import (
	. "github.com/roobert/golang-db"
)

func init() {
	OpenDB("data.db")

	schema := `
		id            INTEGER  PRIMARY KEY,
		dev_id        TEXT     NOT NULL,
		timestamp     DATETIME NOT NULL,
		latitude      REAL     NOT NULL,
		longitude     REAL     NOT NULL,
		direction     TEXT     NOT NULL
	`

	CreateTable("data", schema)
}
