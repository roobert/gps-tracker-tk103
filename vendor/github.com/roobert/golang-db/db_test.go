package db

import (
	"os"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

func TestOpenDB(t *testing.T) {
	func() {
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("panic when trying to create DB")
			} else {
				DeleteDB("test.db")

				if _, err := os.Stat("test.db"); os.IsExist(err) {
					t.Errorf("failed to delete db")
				}
			}
		}()

		OpenDB("test.db")

		// write to db so file gets created
		CreateTable("test", "id INTEGER PRIMARY KEY, name TEXT")

		if _, err := os.Stat("test.db"); os.IsNotExist(err) {
			t.Errorf("failed to create db")
		}
	}()
}

func TestDeleteDB(t *testing.T) {
	func() {
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("panic when trying to delete DB")
			}
		}()

		OpenDB("test.db")

		// write to db so file gets created
		CreateTable("test", "id INTEGER PRIMARY KEY, name TEXT")

		if _, err := os.Stat("test.db"); os.IsNotExist(err) {
			t.Errorf("failed to create db")
		}

		DeleteDB("test.db")

		if _, err := os.Stat("test.db"); os.IsExist(err) {
			t.Errorf("failed to delete db")
		}
	}()
}

func TestCreateTable(t *testing.T) {
	func() {
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("panic when trying to create table")
			} else {
				DeleteDB("test.db")
			}
		}()

		OpenDB("test.db")
		CreateTable("test", "id INTEGER PRIMARY KEY, name TEXT")
	}()
}
