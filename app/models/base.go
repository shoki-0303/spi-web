package models

import (
	"database/sql"
	"log"
	"spi-web/config"

	_ "github.com/mattn/go-sqlite3"
)

var Db *sql.DB

func init() {
	var err error
	Db, err = sql.Open(config.Config.DbDriver, config.Config.DbName)
	if err != nil {
		log.Fatalf("Db err=%s", err)
	}

	cmd := `CREATE TABLE IF NOT EXISTS admin_users (
						id integer PRIMARY KEY AUTOINCREMENT,
						name text NOT NULL UNIQUE CHECK (name != ""),
						email text NOT NULL UNIQUE CHECK (email != ""),
						password text NOT NULL CHECK (password != ""),
						admin_level integer default 3)`
	_, err = Db.Exec(cmd)
	if err != nil {
		log.Printf("create admin_users table, err=%s", err)
	}
}
