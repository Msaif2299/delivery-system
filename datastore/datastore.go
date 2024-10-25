package datastore

import (
	"context"
	"database/sql"

	"github.com/gin-gonic/gin"
)

const (
	SQLDataStoreKey   string = "sqldatastorekey"
	NoSQLDataStoreKey string = "nosqldatastorekey"
)

type SQLDataStore interface {
	Query(string, ...any) (*sql.Rows, error)
	QueryRow(string, ...any) *sql.Row
	Exec(string, ...any) (int64, error)
}

func GetSQLDataStore(c *gin.Context) SQLDataStore {
	return c.MustGet(SQLDataStoreKey).(SQLDataStore)
}

func SQLDBProvider(db SQLDataStore) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Set(SQLDataStoreKey, db)
	}
}

type InfluxDataStore interface {
	Get(context.Context, string) ([]interface{}, error)
	WriteSync(context.Context, string, map[string]string, map[string]interface{}) error
	WriteAsync(context.Context, string, map[string]string, map[string]interface{})
}

func GetNoSQLDataStore(c *gin.Context) InfluxDataStore {
	return c.MustGet(NoSQLDataStoreKey).(InfluxDataStore)
}

func NoSQLDBProvider(db InfluxDataStore) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Set(NoSQLDataStoreKey, db)
	}
}
