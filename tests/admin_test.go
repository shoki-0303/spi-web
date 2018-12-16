package tests

import (
	"database/sql"
	"fmt"
	"log"
	"os"
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
						email text NOT NULL UNIQUE CHECK (email != ""),
						password text NOT NULL CHECK (password != ""),
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
		_, err := tx.Exec(cmd, "", testUser01.email, testUser01.password)
		return err
	})
	if err.Error() != "CHECK constraint failed: test_admin_users" {
		t.Error("when admin user have NULL name, error should be returned and execute RollBack")
	}
}

func TestEmailNullAdminUser(t *testing.T) {
	err := helpers.WithTransaction(Db, func(tx *sql.Tx) error {
		cmd := fmt.Sprintf(`INSERT INTO %s (name, email, password) Values (?, ?, ?)`, tableName)
		_, err := tx.Exec(cmd, testUser01.name, "", testUser01.password)
		return err
	})
	if err.Error() != "CHECK constraint failed: test_admin_users" {
		t.Error("when admin user have NULL email, error should be returned and execute RollBack")
	}
}

func TestPasswordNullAdminUser(t *testing.T) {
	err := helpers.WithTransaction(Db, func(tx *sql.Tx) error {
		cmd := fmt.Sprintf(`INSERT INTO %s (name, email, password) Values (?, ?, ?)`, tableName)
		_, err := tx.Exec(cmd, testUser01.name, testUser01.email, "")
		return err
	})
	if err.Error() != "CHECK constraint failed: test_admin_users" {
		t.Error("when admin user have NULL password, error should be returned and execute RollBack")
	}
}

func TestDuplicateEmailAdminUser(t *testing.T) {
	err := helpers.WithTransaction(Db, func(tx *sql.Tx) error {
		cmd := fmt.Sprintf(`INSERT INTO %s (name, email, password) Values (?, ?, ?)`, tableName)
		_, err := tx.Exec(cmd, testUser01.name, testUser01.email, testUser01.password)
		return err
	})

	err = helpers.WithTransaction(Db, func(tx *sql.Tx) error {
		cmd := fmt.Sprintf(`INSERT INTO %s (name, email, password) Values (?, ?, ?)`, tableName)
		_, err := tx.Exec(cmd, "ddd", testUser01.email, "ddddddddddddd")
		return err
	})

	if err.Error() != "UNIQUE constraint failed: test_admin_users.email" {
		t.Error("when user email is duplicated,error should be returned and execute RollBack")
	}
}

func TestMain(m *testing.M) {
	setup()
	defer teardown()
	ret := m.Run()
	os.Exit(ret)
}
