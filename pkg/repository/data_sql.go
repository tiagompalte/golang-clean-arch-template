package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/tiagompalte/golang-clean-arch-template/configs"
	"github.com/tiagompalte/golang-clean-arch-template/pkg/errors"
)

type DataSql struct {
	db *sql.DB
}

func NewDataSqlWithConfig(config configs.ConfigDatabase) DataSqlManager {
	db, err := sql.Open(config.DriverName.String(), config.ConnectionSource)
	if err != nil {
		panic(fmt.Sprintf("error to connect in database: %v", err))
	}
	return NewDataSql(db)
}

func NewDataSql(db *sql.DB) DataSqlManager {
	data := new(DataSql)
	data.db = db
	return data
}

func (d *DataSql) Begin() (TransactionSqlManager, error) {
	tx, err := d.db.Begin()
	if err != nil {
		return nil, errors.Wrap(err)
	}
	return newTransaction(tx), nil
}

func (d *DataSql) Close() (err error) {
	return d.db.Close()
}

func (d *DataSql) IsHealthy(ctx context.Context) (bool, error) {
	err := d.db.PingContext(ctx)
	if err != nil {
		return false, errors.Wrap(err)
	}
	return true, nil
}

func (d *DataSql) Command() ConnectorSql {
	return ConnectorSql{
		Exec: func(ctx context.Context, query string, args ...any) (ResultSql, error) {
			return d.db.ExecContext(ctx, query, args...)
		},
		QueryRow: func(ctx context.Context, query string, args ...any) RowSql {
			return d.db.QueryRowContext(ctx, query, args...)
		},
		Query: func(ctx context.Context, query string, args ...any) (RowsSql, error) {
			return d.db.QueryContext(ctx, query, args...)
		},
		ValidateUpdateResult: func(ctx context.Context, result ResultSql) error {
			return validateUpdateResult(ctx, result)
		},
	}
}
