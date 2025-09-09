package uow

import (
	"context"
	"errors"
	"fmt"

	"github.com/tiagompalte/golang-clean-arch-template/internal/pkg/infra/data"
	"github.com/tiagompalte/golang-clean-arch-template/internal/pkg/infra/mongo"
	"github.com/tiagompalte/golang-clean-arch-template/internal/pkg/infra/sql"
	"github.com/tiagompalte/golang-clean-arch-template/pkg/repository"
)

type Uow[T repository.Connector, U data.RepositoryManager] struct {
	db             repository.DataManager[T]
	tx             repository.TransactionManager[T]
	newRepoManager func(conn T) data.Manager[U]
}

func NewUowSql(db repository.DataSqlManager) Uow[repository.ConnectorSql, data.SqlManager] {
	return Uow[repository.ConnectorSql, data.SqlManager]{
		db: db,
		newRepoManager: func(conn repository.ConnectorSql) data.Manager[data.SqlManager] {
			return sql.NewRepositoryManager(conn)
		},
	}
}

func NewUowMongo(db repository.DataMongoManager) Uow[repository.ConnectorMongo, data.MongoManager] {
	return Uow[repository.ConnectorMongo, data.MongoManager]{
		db: db,
		newRepoManager: func(conn repository.ConnectorMongo) data.Manager[data.MongoManager] {
			return mongo.NewRepositoryManager(conn)
		},
	}
}

func (u *Uow[T, U]) Do(ctx context.Context, fn func(data data.Manager[U]) error) error {
	if u.tx != nil {
		return fmt.Errorf("transaction already started")
	}

	tx, err := u.db.Begin()
	if err != nil {
		return err
	}
	u.tx = tx
	dataManager := u.newRepoManager(u.tx.Command())

	err = fn(dataManager)
	if err != nil {
		errRb := u.Rollback()
		if errRb != nil {
			return fmt.Errorf("original error: %s, rollback error: %s", err.Error(), errRb.Error())
		}
		return err
	}

	return u.CommitOrRollback()
}

func (u *Uow[T, U]) Rollback() error {
	if u.tx == nil {
		return errors.New("no transaction to rollback")
	}

	err := u.tx.Rollback()
	if err != nil {
		return err
	}

	u.tx = nil

	return nil
}

func (u *Uow[T, U]) CommitOrRollback() error {
	err := u.tx.Commit()
	if err != nil {
		errRb := u.Rollback()
		if errRb != nil {
			return fmt.Errorf("original error: %s, rollback error: %s", err.Error(), errRb.Error())
		}
		return err
	}

	u.tx = nil

	return nil
}
