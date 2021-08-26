package database

import "database/sql"

// Executor defines the interface required for database connections
// or transactions for querying the database
type Executor interface {
	Exec(string, ...interface{}) (sql.Result, error)
	QueryRow(string, ...interface{}) *sql.Row
	Query(string, ...interface{}) (*sql.Rows, error)
	Prepare(query string) (*sql.Stmt, error)
}
