package uow

import (
	"context"
	"errors"
	"fmt"

	"github.com/tiagompalte/golang-clean-arch-template/internal/pkg/infra/data"
	"github.com/tiagompalte/golang-clean-arch-template/pkg/repository"
)

type Uow struct {
	db         repository.DataSqlManager
	tx         repository.TransactionSqlManager
	repository data.RepositoryManager
}

func NewUow(db repository.DataSqlManager) Uow {
	return Uow{
		db: db,
	}
}

func (u *Uow) Repository() data.RepositoryManager {
	return u.repository
}

func (u *Uow) Do(ctx context.Context, fn func(Uow *Uow) error) error {
	if u.tx != nil {
		return fmt.Errorf("transaction already started")
	}

	tx, err := u.db.Begin()
	if err != nil {
		return err
	}
	u.tx = tx
	u.repository = data.NewRepositoryManager(tx.Command())

	err = fn(u)
	if err != nil {
		errRb := u.Rollback()
		if errRb != nil {
			return fmt.Errorf("original error: %s, rollback error: %s", err.Error(), errRb.Error())
		}
		return err
	}

	return u.CommitOrRollback()
}

func (u *Uow) Rollback() error {
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

func (u *Uow) CommitOrRollback() error {
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
