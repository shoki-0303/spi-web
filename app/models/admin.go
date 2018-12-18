package models

import (
	"database/sql"
	"fmt"
	"log"
	"spi-web/app/models/helpers"
)

type AdminUser struct {
	Name           string
	Email          string
	HashedPassword string
}

var tablename = "admin_users"

func CreateAdminUser(adminUser *AdminUser) error {
	err := helpers.WithTransaction(Db, func(tx *sql.Tx) error {
		cmd := fmt.Sprintf(`INSERT INTO %s (name, email, password) VALUES (?, ?, ?)`, tablename)
		_, err := tx.Exec(cmd, adminUser.Name, adminUser.Email, adminUser.HashedPassword)
		if err != nil {
			log.Printf("action=CreateAdminUser err=%s", err)
		}
		return err
	})
	return err
}
