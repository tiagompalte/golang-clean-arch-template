package repository

import (
	"context"
	"database/sql"
)

type Scanner interface {
	Scan(dest ...interface{}) error
}

type Connector interface {
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row
	QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)
}

type DataManager interface {
	Connector
	Begin() (TransactionManager, error)
	Close() error
	PingContext(ctx context.Context) error
}

type TransactionManager interface {
	Connector
	Rollback() error
	Commit() error
}
