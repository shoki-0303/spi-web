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
}
