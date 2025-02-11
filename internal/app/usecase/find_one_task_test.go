package usecase

import (
	"context"
	"testing"
	"time"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"github.com/tiagompalte/golang-clean-arch-template/internal/pkg/infra/data"
	"github.com/tiagompalte/golang-clean-arch-template/pkg/repository"
)

func TestFindOneTaskExecute(t *testing.T) {
	t.Parallel()

	t.Run("should return one task", func(t *testing.T) {
		t.Parallel()
		ctx := context.Background()

		db, mock, err := sqlmock.New()
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}

		datasql := repository.NewDataSql(db)
		taskRepository := data.NewTaskRepository(datasql.Command())

		us := FindOneTaskUseCaseImpl{
			taskRepository: taskRepository,
		}

		mock.ExpectQuery(
			`SELECT (.+) FROM tb_task t JOIN tb_category c ON NOT c.deleted_at AND t.category_id = c.id WHERE NOT t.deleted_at AND t.uuid = \?`,
		).WithArgs("uuid").WillReturnRows(sqlmock.NewRows([]string{
			"t.id", "t.created_at", "t.updated_at", "t.uuid", "t.name", "t.description", "t.done", "c.id", "c.created_at", "c.updated_at", "c.slug", "c.name", "t.user_id",
		}).AddRow(1, time.Time{}, time.Time{}, "uuid", "task", "new task", false, 1, time.Time{}, time.Time{}, "category", "category", 1))

		var input FindOneTaskInput
		input.TaskUUID = "uuid"
		input.UserID = 1

		task, err := us.Execute(ctx, input)

		if err != nil {
			t.Error(err)
		}

		if task.UUID != "uuid" {
			t.Errorf("task uuid should be uuid but is %s", task.UUID)
		}
	})
}
