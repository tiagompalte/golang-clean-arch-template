package usecase

import (
	"context"
	"testing"
	"time"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"github.com/tiagompalte/golang-clean-arch-template/internal/pkg/infra/data"
	"github.com/tiagompalte/golang-clean-arch-template/pkg/crypto"
	"github.com/tiagompalte/golang-clean-arch-template/pkg/errors"
	"github.com/tiagompalte/golang-clean-arch-template/pkg/repository"
)

func TestCreateUserExecute(t *testing.T) {
	t.Parallel()
	crypto := crypto.NewCryptoMock()

	t.Run("should insert new user", func(t *testing.T) {
		t.Parallel()
		ctx := context.Background()

		db, mock, err := sqlmock.New()
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}

		datasql := repository.NewDataSql(db)
		userRepository := data.NewUserRepository(datasql.Command())

		us := CreateUserUseCaseImpl{
			userRepository: userRepository,
			crypto:         crypto,
		}

		mock.ExpectExec(
			`INSERT INTO tb_user \(uuid, name, email, pass_encrypted\) VALUES \(\?,\?,\?,\?\)`,
		).WithArgs(sqlmock.AnyArg(), "User", "user@email.com", "pass").WillReturnResult(sqlmock.NewResult(1, 1))

		mock.ExpectQuery(
			`SELECT (.+) FROM tb_user u WHERE NOT u.deleted_at AND u.id = \?`,
		).WithArgs(1).WillReturnRows(sqlmock.NewRows([]string{
			"u.id", "u.created_at", "u.updated_at", "u.version", "u.uuid", "u.name", "u.email",
		}).AddRow(1, time.Time{}, time.Time{}, 1, "uuid", "User", "user@email.com"))

		input := CreateUserInput{
			Name:     "User",
			Email:    "user@email.com",
			Password: "pass",
		}
		result, err := us.Execute(ctx, input)

		if err != nil {
			t.Error(err)
		}

		if result.UUID != "uuid" {
			t.Errorf("uuid should be uuid but is %s", result.UUID)
		}

		if result.Version != 1 {
			t.Errorf("version should be 1 but is %d", result.Version)
		}

		if result.Name != "User" {
			t.Errorf("name should be User but is %s", result.Name)
		}

		if result.Email != "user@email.com" {
			t.Errorf("email should be user@email.com but is %s", result.Email)
		}
	})

	t.Run("should return error if repository insert return error", func(t *testing.T) {
		t.Parallel()
		ctx := context.Background()

		db, mock, err := sqlmock.New()
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}

		datasql := repository.NewDataSql(db)
		userRepository := data.NewUserRepository(datasql.Command())

		us := CreateUserUseCaseImpl{
			userRepository: userRepository,
			crypto:         crypto,
		}

		mock.ExpectExec(
			`INSERT INTO tb_user \(uuid, name, email, pass_encrypted\) VALUES \(\?,\?,\?,\?\)`,
		).WithArgs(sqlmock.AnyArg(), "User", "user@email.com", "pass").WillReturnError(errors.NewAppConflictError("user"))

		input := CreateUserInput{
			Name:     "User",
			Email:    "user@email.com",
			Password: "pass",
		}
		result, err := us.Execute(ctx, input)

		if !errors.IsAppError(err, errors.ErrorCodeConflict) {
			t.Errorf("error should be conflict but is %s", errors.Cause(err))
		}

		if result.ID != 0 {
			t.Errorf("id should be 0 but is %d", result.ID)
		}
	})

	t.Run("should return error if is invalid", func(t *testing.T) {
		t.Parallel()
		ctx := context.Background()

		db, _, err := sqlmock.New()
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}

		datasql := repository.NewDataSql(db)
		userRepository := data.NewUserRepository(datasql.Command())

		us := CreateUserUseCaseImpl{
			userRepository: userRepository,
			crypto:         crypto,
		}

		input := CreateUserInput{
			Name:     "",
			Email:    "",
			Password: "",
		}
		_, err = us.Execute(ctx, input)

		if err == nil {
			t.Errorf("should be return error but is nil")
		}
	})
}
