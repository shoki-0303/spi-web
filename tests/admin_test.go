package tests

import (
	"database/sql"
	"fmt"
	"log"
	"spi-web/app/models/helpers"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

var Db *sql.DB

func setup() {
	var err error
	Db, err = sql.Open("sqlite3", "spitest.sql")
	if err != nil {
		log.Printf("action=TestInit err=%s", err)
	}

	cmd := `CREATE TABLE IF NOT EXISTS test_admin_users (
						id integer PRIMARY KEY AUTOINCREMENT,
						name text NOT NULL CHECK (name != ""),
						email text NOT NULL UNIQUE CHECK (name != ""),
						password text NOT NULL CHECK (name != ""),
						admin_level integer default 3)`
	_, err = Db.Exec(cmd)
	if err != nil {
		log.Printf("create test_admin_users table err=%s", err)
	}
	fmt.Println("setup")
}

func teardown() {
	Db.Close()
	fmt.Println("teardown")
}

type TestAdminUser struct {
	name     string
	email    string
	password string
}

var testUser01 = &TestAdminUser{
	name:     "testuser01",
	email:    "testuser01@gmail.com",
	password: "0987654321",
}

var tableName = "test_admin_users"

func TestNameNullAdminUser(t *testing.T) {
	err := helpers.WithTransaction(Db, func(tx *sql.Tx) error {
		cmd := fmt.Sprintf(`INSERT INTO %s (name, email, password) Values (?, ?, ?)`, tableName)
		_, err := tx.Exec(cmd, "", "aaa@gmail.com", "qwrjbjasbdjasd")
		return err
	})
	if err.Error() != "CHECK constraint failed: test_admin_users" {
		t.Error("when admin user have NULL name, error should be returned and execute RollBack")
	}
}

func TestMain(m *testing.M) {
	setup()
	m.Run()
	defer teardown()
}
