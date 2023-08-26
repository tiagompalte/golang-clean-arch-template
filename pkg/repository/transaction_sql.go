package repository

import (
	"context"
	"database/sql"

	"github.com/tiagompalte/golang-clean-arch-template/pkg/errors"
)

type TransactionSql struct {
	tx         *sql.Tx
	committed  bool
	rolledback bool
}

func newTransaction(tx *sql.Tx) TransactionManager {
	transaction := new(TransactionSql)
	transaction.tx = tx
	return transaction
}

func (t *TransactionSql) ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	return t.tx.ExecContext(ctx, query, args...)
}

func (t *TransactionSql) QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row {
	return t.tx.QueryRowContext(ctx, query, args...)
}

func (t *TransactionSql) QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error) {
	return t.tx.QueryContext(ctx, query, args...)
}

func (t *TransactionSql) Commit() error {
	err := t.tx.Commit()
	if err != nil {
		return errors.Wrap(err)
	}

	t.committed = true

	return nil
}

func (t *TransactionSql) Rollback() error {
	if t != nil && !t.committed && !t.rolledback {
		err := t.tx.Rollback()
		if err != nil {
			return errors.Wrap(err)
		}
		t.rolledback = true
	}
	return nil
}
