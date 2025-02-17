package repository

import (
	"context"

	"github.com/tiagompalte/golang-clean-arch-template/pkg/healthcheck"
)

type Scanner interface {
	Scan(dest ...any) error
}

type ResultSql interface {
	LastInsertId() (int64, error)
	RowsAffected() (int64, error)
}

type RowSql interface {
	Scanner
	Err() error
}

type RowsSql interface {
	Scanner
	Next() bool
	Close() error
}

type ConnectorSql struct {
	Exec                 func(ctx context.Context, query string, args ...any) (ResultSql, error)
	QueryRow             func(ctx context.Context, query string, args ...any) RowSql
	Query                func(ctx context.Context, query string, args ...any) (RowsSql, error)
	ValidateUpdateResult func(ctx context.Context, result ResultSql) error
}

type ConnectorMongo struct {
	InsertOne func(ctx context.Context, doc any) error
}

type Connector interface {
	ConnectorSql | ConnectorMongo
}

type DataSqlManager = DataManager[ConnectorSql]

type TransactionSqlManager = TransactionManager[ConnectorSql]

type DataMongoManager = DataManager[ConnectorMongo]

type TransactionMongoManager = TransactionManager[ConnectorMongo]

type DataManager[T Connector] interface {
	healthcheck.HealthCheck
	Command() T
	Begin() (TransactionManager[T], error)
	Close() error
}

type TransactionManager[T Connector] interface {
	Command() T
	Rollback() error
	Commit() error
}
