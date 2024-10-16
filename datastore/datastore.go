package datastore

import (
	"database/sql"

	"github.com/gin-gonic/gin"
)

const SQLDataStoreKey string = "sqldatastorekey"

type SQLDataStore interface {
	Query(string, ...any) (*sql.Rows, error)
	QueryRow(string, ...any) *sql.Row
}

func GetSQLDataStore(c *gin.Context) SQLDataStore {
	return c.MustGet(SQLDataStoreKey).(SQLDataStore)
}

func SQLDBProvider(db SQLDataStore) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Set(SQLDataStoreKey, db)
	}
}
