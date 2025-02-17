package usecase

import (
	"context"
	"database/sql"
	"slices"
	"testing"
	"time"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"github.com/tiagompalte/golang-clean-arch-template/internal/pkg/infra/data"
	"github.com/tiagompalte/golang-clean-arch-template/pkg/errors"
	"github.com/tiagompalte/golang-clean-arch-template/pkg/repository"
)

func TestUpdateUserNameUseCaseExecute(t *testing.T) {
	t.Parallel()

	t.Run("Should return error when input is invalid", func(t *testing.T) {
		t.Parallel()
		ctx := context.Background()

		us := UpdateUserNameUseCaseImpl{}

		_, err := us.Execute(ctx, UpdateUserNameInput{
			UserUUID:    "",
			NewUserName: "",
		})

		if err == nil {
			t.Errorf("should be return error but is nil")
		}

		originalErr := errors.Cause(err)
		aggrErr, ok := originalErr.(errors.AggregatedError)
		if !ok {
			t.Errorf("it is not aggregated error type")
		}

		for _, err := range aggrErr {
			appErr, ok := err.(errors.AppError)
			if !ok {
				t.Errorf("it is not app error type")
			}
			if appErr.Code != errors.ErrorCodeEmptyParameter {
				t.Errorf("code is not %s", errors.ErrorCodeEmptyParameter)
			}
			if !slices.Contains([]string{"user_uuid", "user_name"}, appErr.Field) {
				t.Errorf("field is not incorret")
			}
		}
	})

	t.Run("Should update user name", func(t *testing.T) {
		t.Parallel()
		ctx := context.Background()

		db, mock, err := sqlmock.New()
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}

		datasql := repository.NewDataSql(db)
		userRepository := data.NewUserRepository(datasql.Command())

		us := UpdateUserNameUseCaseImpl{
			userRepository: userRepository,
		}

		userID := uint32(1)
		userUUID := "user_uuid"
		newUserName := "new_user_name"
		userEmail := "user@email.com"
		version := uint32(1)

		// FindByUUID
		mock.ExpectQuery(
			`SELECT (.+) FROM tb_user u WHERE NOT u.deleted_at AND u.uuid = \?`,
		).WithArgs(userUUID).WillReturnRows(sqlmock.NewRows([]string{
			"u.id", "u.created_at", "u.updated_at", "u.version", "u.uuid", "u.name", "u.email",
		}).AddRow(userID, time.Time{}, time.Time{}, 1, userUUID, "User", userEmail))

		// UpdateName
		mock.ExpectExec(
			`UPDATE tb_user SET name = \? , version = version \+ 1 WHERE NOT deleted_at AND id = \? AND version = \?`,
		).WithArgs(
			newUserName,
			userID,
			version,
		).WillReturnResult(sqlmock.NewResult(1, 1))

		// FindByID
		mock.ExpectQuery(
			`SELECT (.+) FROM tb_user u WHERE NOT u.deleted_at AND u.id = \?`,
		).WithArgs(userID).WillReturnRows(sqlmock.NewRows([]string{
			"u.id", "u.created_at", "u.updated_at", "u.version", "u.uuid", "u.name", "u.email",
		}).AddRow(userID, time.Time{}, time.Time{}, 2, userUUID, newUserName, userEmail))

		input := UpdateUserNameInput{
			UserUUID:    userUUID,
			NewUserName: newUserName,
			Version:     version,
		}

		result, err := us.Execute(ctx, input)

		if err != nil {
			t.Errorf("should not return error but return %v", err)
		}

		if result.UUID != userUUID {
			t.Errorf("uuid should equal %s but is %s", userUUID, result.UUID)
		}
		if result.Name != newUserName {
			t.Errorf("name should equal %s but is %s", newUserName, result.Name)
		}
		if result.Email != userEmail {
			t.Errorf("email should equal %s but is %s", userEmail, result.Email)
		}
		if result.Version != 2 {
			t.Errorf("version should equal 1 but is %d", result.Version)
		}
	})

	t.Run("Should return error when version is not equal", func(t *testing.T) {
		t.Parallel()
		ctx := context.Background()

		db, mock, err := sqlmock.New()
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}

		datasql := repository.NewDataSql(db)
		userRepository := data.NewUserRepository(datasql.Command())

		us := UpdateUserNameUseCaseImpl{
			userRepository: userRepository,
		}

		userID := uint32(1)
		userUUID := "user_uuid"
		newUserName := "new_user_name"
		userEmail := "user@email.com"
		version := uint32(0)

		// FindByUUID
		mock.ExpectQuery(
			`SELECT (.+) FROM tb_user u WHERE NOT u.deleted_at AND u.uuid = \?`,
		).WithArgs(userUUID).WillReturnRows(sqlmock.NewRows([]string{
			"u.id", "u.created_at", "u.updated_at", "u.version", "u.uuid", "u.name", "u.email",
		}).AddRow(userID, time.Time{}, time.Time{}, 1, userUUID, "User", userEmail))

		// UpdateName
		mock.ExpectExec(
			`UPDATE tb_user SET name = \? , version = version \+ 1 WHERE NOT deleted_at AND id = \? AND version = \?`,
		).WithArgs(
			newUserName,
			userID,
			version,
		).WillReturnResult(sqlmock.NewResult(0, 0))

		input := UpdateUserNameInput{
			UserUUID:    userUUID,
			NewUserName: newUserName,
			Version:     version,
		}

		_, err = us.Execute(ctx, input)

		if err == nil {
			t.Errorf("should return error but is nil")
		}

		errApp, ok := errors.Cause(err).(errors.AppError)
		if !ok {
			t.Errorf("error is not app error type")
		}
		if errApp.Code != errors.ErrorCodeConcurrencyRepository {
			t.Errorf("error code is not %s", errors.ErrorCodeConcurrencyRepository)
		}
	})

	t.Run("Should return error when user is not found", func(t *testing.T) {
		t.Parallel()
		ctx := context.Background()

		db, mock, err := sqlmock.New()
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}

		datasql := repository.NewDataSql(db)
		userRepository := data.NewUserRepository(datasql.Command())

		us := UpdateUserNameUseCaseImpl{
			userRepository: userRepository,
		}

		userUUID := "user_uuid"
		newUserName := "new_user_name"
		version := uint32(0)

		// FindByUUID
		mock.ExpectQuery(
			`SELECT (.+) FROM tb_user u WHERE NOT u.deleted_at AND u.uuid = \?`,
		).WithArgs(userUUID).WillReturnError(sql.ErrNoRows)

		input := UpdateUserNameInput{
			UserUUID:    userUUID,
			NewUserName: newUserName,
			Version:     version,
		}

		_, err = us.Execute(ctx, input)

		if err == nil {
			t.Errorf("should return error but is nil")
		}

		errApp, ok := errors.Cause(err).(errors.AppError)
		if !ok {
			t.Errorf("error is not app error type")
		}
		if errApp.Code != errors.ErrorCodeNotFound {
			t.Errorf("error code is not %s", errors.ErrorCodeNotFound)
		}
	})
}
