package tests

import (
	"database/sql"
	"fmt"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gin-gonic/gin"
)

type EndPointType string

const (
	GET    EndPointType = "GET"
	PUT    EndPointType = "PUT"
	DELETE EndPointType = "DELETE"
	POST   EndPointType = "POST"
)

type MockSQLDataStore struct {
	conn *sql.DB
}

func (db *MockSQLDataStore) Query(query string, params ...any) (*sql.Rows, error) {
	rows, err := db.conn.Query(query, params...)
	if err != nil {
		return nil, err
	}
	return rows, nil
}

func (db *MockSQLDataStore) QueryRow(query string, params ...any) *sql.Row {
	row := db.conn.QueryRow(query, params...)
	return row
}

func (db *MockSQLDataStore) Exec(query string, params ...any) (int64, error) {
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

func mockSQLDBProvider() (gin.HandlerFunc, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		panic(fmt.Sprintf("Error encountered while making mock, err: %s", err.Error()))
	}
	return func(ctx *gin.Context) {
		ctx.Set("sqldatastorekey", &MockSQLDataStore{conn: db})
		ctx.Next()
	}, mock
}

func setupRouter(
	parentGroup string,
	endPoint string,
	endPointType EndPointType,
	handler func(*gin.Context),
) (*gin.Engine, sqlmock.Sqlmock) {
	r := gin.Default()
	dbProvider, mock := mockSQLDBProvider()
	group := r.Group(parentGroup, dbProvider)
	{
		switch endPointType {
		case GET:
			group.GET(endPoint, handler)
		case POST:
			group.POST(endPoint, handler)
		case PUT:
			group.PUT(endPoint, handler)
		case DELETE:
			group.DELETE(endPoint, handler)
		}
	}
	return r, mock
}
