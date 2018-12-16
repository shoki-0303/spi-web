package models

import (
	"database/sql"
	"fmt"
	"log"
	"spi-web/app/models/helpers"
)

var tablename = "admin_users"

func CreateAdminUser() error {
	err := helpers.WithTransaction(Db, func(tx *sql.Tx) error {
		cmd := fmt.Sprintf(`INSERT INTO %s (name, email, password) VALUES (?, ?, ?)`, tablename)
		_, err := tx.Exec(cmd, "user001", "user001@gmail.com", "0987654321")
		if err != nil {
			log.Printf("action=CreateAdminUser err=%s", err)
		}
		return err
	})

	return err
}
