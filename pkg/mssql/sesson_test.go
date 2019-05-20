package mssql_test

import (
	"database/sql"
	"b2yun/pkg/mssql"

	"github.com/jmoiron/sqlx"
	sqlmock "gopkg.in/DATA-DOG/go-sqlmock.v1"
)

type MockSession struct {
	mockDB       *sql.DB
	sqlxDB       *sqlx.DB
	mock         sqlmock.Sqlmock
	mssqlSession *mssql.Session
}

func NewMockSession() (MockSession, error) {

	var mockSession MockSession

	mockDB, mock, err := sqlmock.New()
	if err != nil {
		return mockSession, err
	}

	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")

	mssqlSession := mssql.NewSession(sqlxDB)

	mockSession.sqlxDB = sqlxDB
	mockSession.mock = mock
	mockSession.mockDB = mockDB
	mockSession.mssqlSession = mssqlSession

	return mockSession, nil

}
