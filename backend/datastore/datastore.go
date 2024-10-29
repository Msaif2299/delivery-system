package datastore

import (
	"context"

	"github.com/gin-gonic/gin"
)

const (
	SQLDataStoreKey   string = "sqldatastorekey"
	NoSQLDataStoreKey string = "nosqldatastorekey"
)

type SQLDataStore interface {
	Select(ctx context.Context, query string, cacheKey string, dest []interface{}, params ...any) error
	Get(ctx context.Context, query string, cacheKey string, dest interface{}, params ...any) error
	Exec(query string, params ...any) (int64, error)
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
	GetLastValue(ctx context.Context, measurement string, tags map[string]string) (interface{}, error)
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
