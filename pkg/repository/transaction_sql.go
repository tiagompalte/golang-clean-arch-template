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

func newTransaction(tx *sql.Tx) TransactionSqlManager {
	transaction := new(TransactionSql)
	transaction.tx = tx
	return transaction
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

func (t *TransactionSql) Command() ConnectorSql {
	return ConnectorSql{
		Exec: func(ctx context.Context, query string, args ...any) (ResultSql, error) {
			return t.tx.ExecContext(ctx, query, args...)
		},
		QueryRow: func(ctx context.Context, query string, args ...any) RowSql {
			return t.tx.QueryRowContext(ctx, query, args...)
		},
		Query: func(ctx context.Context, query string, args ...any) (RowsSql, error) {
			return t.tx.QueryContext(ctx, query, args...)
		},
		ValidateUpdateResult: func(ctx context.Context, result ResultSql) error {
			return validateUpdateResult(ctx, result)
		},
	}
}
