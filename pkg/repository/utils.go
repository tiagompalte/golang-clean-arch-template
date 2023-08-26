package repository

import (
	"database/sql"
	"strings"

	"github.com/tiagompalte/golang-clean-arch-template/pkg/errors"
)

const (
	sqlNull  = "NULL"
	sqlEmpty = "''"
)

func FmtParamList(length int, or ...string) string {
	opt := sqlNull
	if len(or) != 0 {
		opt = or[0]
	}
	if length == 0 {
		return opt
	}
	return strings.Repeat(",?", length)[1:]
}

func ParseEntities[T any](scan func(Scanner) (T, error), rows *sql.Rows, err error) ([]T, error) {
	if err != nil {
		return nil, errors.Wrap(err)
	}

	defer rows.Close()

	entitySet := make([]T, 0)
	for rows.Next() {
		entity, err := scan(rows)
		if err != nil {
			return nil, errors.Wrap(err)
		}
		entitySet = append(entitySet, entity)
	}

	return entitySet, nil
}
