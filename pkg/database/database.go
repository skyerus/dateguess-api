package database

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// New sets up a DB connection and returns an sql.DB pointer
func New(config Config) (*sql.DB, error) {
	db, err := sql.Open("mysql", config.ConnectionString())
	if err != nil {
		return nil, err
	}

	db.SetMaxIdleConns(0)

	return db, nil
}

// ScopedTx provide a simple way to wrap transactions.
// Automatically deals with commit/rollback based on whether the called
// function returns an error or panics.
func ScopedTx(
	ctx context.Context,
	db *sql.DB,
	opts *sql.TxOptions,
	f func(tx *sql.Tx) error,
) (err error) {
	tx, err := db.BeginTx(ctx, opts)
	if err != nil {
		return err
	}

	defer func() {
		//nolint:gocritic
		if p := recover(); p != nil {
			err := tx.Rollback()
			if err != nil {
				log.Println(err)
			}
			panic(p) // re-throw the panic
		} else if err != nil {
			e := tx.Rollback()
			if e != nil {
				err = fmt.Errorf("%s: %w", e, err)
			}
		} else {
			err = tx.Commit()
		}
	}()

	// NOTE: we do this err = stuff so that err is set for the defer block
	// where we decide whether to commit or rollback.
	err = f(tx)
	return err
}
