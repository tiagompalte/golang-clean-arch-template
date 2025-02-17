package usecase

import (
	"context"
	"testing"
	"time"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"github.com/tiagompalte/golang-clean-arch-template/internal/pkg/infra/data"
	"github.com/tiagompalte/golang-clean-arch-template/pkg/repository"
)

func TestFindUserUUIDExecute(t *testing.T) {
	t.Parallel()

	t.Run("should be find user by UUID", func(t *testing.T) {
		t.Parallel()
		ctx := context.Background()

		db, mock, err := sqlmock.New()
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}

		datasql := repository.NewDataSql(db)
		userRepository := data.NewUserRepository(datasql.Command())

		us := FindUserUUIDUseCaseImpl{
			userRepository: userRepository,
		}

		mock.ExpectQuery(
			`SELECT (.+) FROM tb_user u WHERE NOT u.deleted_at AND u.uuid = \?`,
		).WithArgs("uuid").WillReturnRows(sqlmock.NewRows([]string{
			"u.id", "u.created_at", "u.updated_at", "u.version", "u.uuid", "u.name", "u.email",
		}).AddRow(1, time.Time{}, time.Time{}, 1, "uuid", "User", "user@email.com"))

		user, err := us.Execute(ctx, "uuid")

		if err != nil {
			t.Error(err)
		}

		if user.UUID != "uuid" {
			t.Errorf("task uuid should be uuid but is %s", user.UUID)
		}

		if user.Version != 1 {
			t.Errorf("version should be 1 but is %d", user.Version)
		}
	})
}
