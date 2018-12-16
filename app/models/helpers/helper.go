package helpers

import (
	"database/sql"
	"log"
)

type TxFn func(*sql.Tx) error

func WithTransaction(Db *sql.DB, fn TxFn) error {
	tx, err := Db.Begin()
	if err != nil {
		log.Printf("action=WithTransaction err=%s", err)
		return err
	}

	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
		} else if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	err = fn(tx)
	return err
}
