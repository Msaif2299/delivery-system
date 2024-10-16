package datastore

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	"github.com/go-sql-driver/mysql"
)

type MySQLDataStore struct {
	conn *sql.DB
}

func NewMySQLDataStore() (*MySQLDataStore, error) {
	cfg := mysql.Config{
		User:                    os.Getenv("MYSQL_USER"),
		Passwd:                  os.Getenv("MYSQL_PASSWORD"),
		Net:                     "tcp",
		Addr:                    "mysql",
		DBName:                  os.Getenv("MYSQL_DATABASE"),
		AllowNativePasswords:    true,
		AllowCleartextPasswords: true,
	}
	conn, err := sql.Open("mysql", cfg.FormatDSN())
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
		conn: conn,
	}, nil
}

func (db *MySQLDataStore) Query(query string) (*sql.Rows, error) {
	rows, err := db.conn.Query(query)
	if err != nil {
		return nil, err
	}
	return rows, nil
}

func (db *MySQLDataStore) QueryRow(query string) *sql.Row {
	row := db.conn.QueryRow(query)
	return row
}
