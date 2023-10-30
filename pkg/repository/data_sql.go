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

func NewDataSqlWithConfig(config configs.ConfigDatabase) DataManager {
	db, err := sql.Open(config.DriverName, config.ConnectionSource)
	if err != nil {
		panic(fmt.Sprintf("error to connect in database: %v", err))
	}
	return NewDataSql(db)
}

func NewDataSql(db *sql.DB) DataManager {
	data := new(DataSql)
	data.db = db
	return data
}

func (d *DataSql) Begin() (TransactionManager, error) {
	tx, err := d.db.Begin()
	if err != nil {
		return nil, errors.Wrap(err)
	}
	return newTransaction(tx), nil
}

func (d *DataSql) Close() (err error) {
	return d.db.Close()
}

func (d *DataSql) ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	return d.db.ExecContext(ctx, query, args...)
}

func (d *DataSql) QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row {
	return d.db.QueryRowContext(ctx, query, args...)
}

func (d *DataSql) QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error) {
	return d.db.QueryContext(ctx, query, args...)
}

func (d *DataSql) PingContext(ctx context.Context) error {
	return d.db.PingContext(ctx)
}
