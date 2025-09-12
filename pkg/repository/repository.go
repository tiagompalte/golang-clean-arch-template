package repository

import (
	"context"

	"github.com/tiagompalte/golang-clean-arch-template/pkg/healthcheck"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
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

type RowsMongo interface {
	Next(ctx context.Context) bool
	Decode(value any) error
}

type ResultInsertMongo struct {
	InsertedID any
}

type ResultInsertManyMongo struct {
	InsertedIDs []any
}

type ResultUpdateMongo struct {
	AffectedCount int64
}

type ResultDeleteMongo struct {
	AffectedCount int64
}

type ConnectorMongo struct {
	Aggregate  func(ctx context.Context, collection string, pipeline any) (RowsMongo, error)
	Find       func(ctx context.Context, collection string, filter any, opts ...options.Lister[options.FindOptions]) (RowsMongo, error)
	FindOne    func(ctx context.Context, collection string, filter any, result any) error
	InsertOne  func(ctx context.Context, collection string, doc any) (ResultInsertMongo, error)
	InsertMany func(ctx context.Context, collection string, docs []any) (ResultInsertManyMongo, error)
	UpdateOne  func(ctx context.Context, collection string, filter any, update any) (ResultUpdateMongo, error)
	UpdateMany func(ctx context.Context, collection string, filter any, update any) (ResultUpdateMongo, error)
	DeleteOne  func(ctx context.Context, collection string, filter any) (ResultDeleteMongo, error)
	DeleteMany func(ctx context.Context, collection string, filter any) (ResultDeleteMongo, error)
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
