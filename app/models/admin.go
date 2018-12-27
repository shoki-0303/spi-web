package models

import (
	"database/sql"
	"fmt"
	"log"
	"spi-web/app/models/helpers"

	"golang.org/x/crypto/bcrypt"
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

func GetAdminUser(name string) (*AdminUser, error) {
	var adminUser AdminUser
	err := helpers.WithTransaction(Db, func(tx *sql.Tx) error {
		cmd := fmt.Sprintf(`SELECT name FROM %s WHERE name = ?`, tablename)
		row := tx.QueryRow(cmd, name)
		err := row.Scan(&adminUser.Name)
		if err != nil {
			log.Printf("action=GetAdminUser err=%s", err)
		}
		return err
	})
	return &adminUser, err
}

func ConfirmAdminUser(email, password string) (bool, error, AdminUser) {
	var adminUser AdminUser
	helpers.WithTransaction(Db, func(tx *sql.Tx) error {
		cmd := fmt.Sprintf(`SELECT name, email, password FROM %s WHERE email = ?`, tablename)
		row := tx.QueryRow(cmd, email)
		err := row.Scan(&adminUser.Name, &adminUser.Email, &adminUser.HashedPassword)
		if err != nil {
			log.Printf("action=confirmAdminUser err=%s", err)
		}
		return err
	})

	err := bcrypt.CompareHashAndPassword([]byte(adminUser.HashedPassword), []byte(password))
	if err == nil {
		return true, nil, adminUser
	}
	return false, err, adminUser
}

func UpdateAdminUser(rename, oldname string) (AdminUser, error) {
	var adminUser AdminUser
	err := helpers.WithTransaction(Db, func(*sql.Tx) error {
		cmd := fmt.Sprintf(`UPDATE %s SET name = ? WHERE name = ?`, tablename)
		_, err := Db.Exec(cmd, rename, oldname)
		if err != nil {
			log.Printf("action=UpdateAdminUser, err=%s", err)
		}
		cmd = fmt.Sprintf(`SELECT name, email, password from %s WHERE name = ?`, tablename)
		row := Db.QueryRow(cmd, rename)
		err = row.Scan(&adminUser.Name, &adminUser.Email, &adminUser.HashedPassword)
		if err != nil {
			log.Printf("action=UpdateAdminUser err=%s", err)
		}
		return err
	})
	return adminUser, err
}
