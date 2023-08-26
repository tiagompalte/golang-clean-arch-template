package errors

import (
	"database/sql"
	"strings"

	"github.com/go-sql-driver/mysql"
)

const conflictErrorNumber uint16 = 1062

func Repo(err error, mainTable string) error {
	if err == nil {
		return nil
	}

	entity := strings.TrimPrefix(mainTable, "tb_")
	originalErr := Cause(err)

	switch originalErr {
	case sql.ErrNoRows:
		return Wrap(NewAppNotFoundError(entity), err.Error())
	}

	mysqlErr, ok := originalErr.(*mysql.MySQLError)
	if ok {
		switch mysqlErr.Number {
		case conflictErrorNumber:
			return Wrap(NewAppConflictError(entity), err.Error())
		}
	}

	return Wrap(err)
}
