package datastore

import "database/sql"

type SQLDataStore interface {
	Query(string) (*sql.Rows, error)
	QueryRow(string) *sql.Row
}
