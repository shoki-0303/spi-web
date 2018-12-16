package tests

import (
	"database/sql"
	"fmt"
	"log"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func setup() {
	var err error
	db, err = sql.Open("sqlite3", "spitest.sql")
	if err != nil {
		log.Printf("action=TestInit err=%s", err)
	}

	cmd := `CREATE TABLE IF NOT EXISTS test_admin_users (
						id integer PRIMARY KEY AUTOINCREMENT,
						name text NOT NULL,
						email text NOT NULL UNIQUE,
						password text NOT NULL,
						admin_level integer default 3)`
	_, err = db.Exec(cmd)
	if err != nil {
		log.Printf("create test_admin_users table err=%s", err)
	}
	fmt.Println("setup")
}

func teardown() {
	db.Close()
	fmt.Println("teardown")
}

func TestMain(m *testing.M) {
	setup()
	m.Run()
	defer teardown()
}
