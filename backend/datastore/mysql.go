package datastore

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type MySQLDataStore struct {
	conn  *sqlx.DB
	cache Cache
}

func NewMySQLDataStore(cache Cache) (*MySQLDataStore, error) {
	cfg := mysql.Config{
		User:                    os.Getenv("MYSQL_USER"),
		Passwd:                  os.Getenv("MYSQL_PASSWORD"),
		Net:                     "tcp",
		Addr:                    "mysql",
		DBName:                  os.Getenv("MYSQL_DATABASE"),
		AllowNativePasswords:    true,
		AllowCleartextPasswords: true,
	}
	conn, err := sqlx.Open("mysql", cfg.FormatDSN())
	if err != nil {
		return nil, fmt.Errorf("unable to connect to sql server, err: %s", err.Error())
	}
	for retryCount := 0; retryCount < 10; retryCount++ {
		pingErr := conn.Ping()
		if pingErr != nil {
			fmt.Printf("Unable to send pings to sql server, err: %s,  retrying in %d seconds\n", pingErr.Error(), 2*retryCount+1)
			time.Sleep(time.Duration(2*retryCount+1) * time.Second)
			if retryCount == 2*10-1 {
				return nil, fmt.Errorf("unable to send pings to sql server, err: %s", pingErr.Error())
			}
			continue
		}
		break
	}
	return &MySQLDataStore{
		conn:  conn,
		cache: cache,
	}, nil
}

func (db *MySQLDataStore) Select(ctx context.Context, query string, cacheKey string, dest []interface{}, params ...any) error {
	if err := db.cache.Get(ctx, cacheKey, dest); err == nil {
		return nil
	}
	err := db.conn.Select(dest, query, params...)
	return err
}

func (db *MySQLDataStore) Get(ctx context.Context, query string, cacheKey string, dest interface{}, params ...any) error {
	if err := db.cache.Get(ctx, cacheKey, dest); err == nil {
		return nil
	}
	err := db.conn.Get(dest, query, params...)
	return err
}

func (db *MySQLDataStore) Exec(query string, params ...any) (int64, error) {
	result, err := db.conn.Exec(query, params...)
	if err != nil {
		return 0, err
	}
	id, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}
	return id, nil
}
