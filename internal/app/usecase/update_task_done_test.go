package usecase

import (
	"context"
	"testing"
	"time"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"github.com/tiagompalte/golang-clean-arch-template/internal/pkg/infra/data"
	"github.com/tiagompalte/golang-clean-arch-template/pkg/errors"
	"github.com/tiagompalte/golang-clean-arch-template/pkg/repository"
)

func TestUpdateTaskDoneExecute(t *testing.T) {
	t.Parallel()

	t.Run("should be update task to done", func(t *testing.T) {
		t.Parallel()
		ctx := context.Background()

		db, mock, err := sqlmock.New()
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}

		datasql := repository.NewDataSql(db)
		taskRepository := data.NewTaskRepository(datasql.Command())

		us := UpdateTaskDoneUseCaseImpl{
			taskRepository: taskRepository,
		}

		mock.ExpectQuery(
			`SELECT (.+) FROM tb_task t JOIN tb_category c ON NOT c.deleted_at AND t.category_id = c.id WHERE NOT t.deleted_at AND t.uuid = \?`,
		).WithArgs("uuid").WillReturnRows(sqlmock.NewRows([]string{
			"t.id", "t.created_at", "t.updated_at", "t.uuid", "t.name", "t.description", "t.done", "c.id", "c.created_at", "c.updated_at", "c.slug", "c.name", "t.user_id",
		}).AddRow(1, time.Time{}, time.Time{}, "uuid", "task", "new task", false, 1, time.Time{}, time.Time{}, "category", "category", 1))

		mock.ExpectExec(`UPDATE tb_task SET done = \? WHERE NOT deleted_at AND id = \?`).WithArgs(true, 1).WillReturnResult(sqlmock.NewResult(1, 1))

		err = us.Execute(ctx, UpdateTaskDoneUseCaseInput{
			UserID: uint32(1),
			UUID:   "uuid",
		})

		if err != nil {
			t.Error(err)
		}
	})

	t.Run("should be return invalid user error if userID is differente", func(t *testing.T) {
		t.Parallel()
		ctx := context.Background()

		db, mock, err := sqlmock.New()
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}

		datasql := repository.NewDataSql(db)
		taskRepository := data.NewTaskRepository(datasql.Command())

		us := UpdateTaskDoneUseCaseImpl{
			taskRepository: taskRepository,
		}

		mock.ExpectQuery(
			`SELECT (.+) FROM tb_task t JOIN tb_category c ON NOT c.deleted_at AND t.category_id = c.id WHERE NOT t.deleted_at AND t.uuid = \?`,
		).WithArgs("uuid").WillReturnRows(sqlmock.NewRows([]string{
			"t.id", "t.created_at", "t.updated_at", "t.uuid", "t.name", "t.description", "t.done", "c.id", "c.created_at", "c.updated_at", "c.slug", "c.name", "t.user_id",
		}).AddRow(1, time.Time{}, time.Time{}, "uuid", "task", "new task", false, 1, time.Time{}, time.Time{}, "category", "category", 1))

		err = us.Execute(ctx, UpdateTaskDoneUseCaseInput{
			UserID: uint32(2),
			UUID:   "uuid",
		})

		if !errors.IsAppError(err, errors.ErrorCodeInvalidUser) {
			t.Error(err)
		}
	})
}
