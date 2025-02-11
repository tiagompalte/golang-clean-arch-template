package usecase

import (
	"context"
	"testing"
	"time"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"github.com/tiagompalte/golang-clean-arch-template/internal/pkg/infra/data"
	"github.com/tiagompalte/golang-clean-arch-template/pkg/repository"
)

func TestFindAllTaskExecute(t *testing.T) {
	t.Parallel()

	t.Run("should return list tasks", func(t *testing.T) {
		t.Parallel()
		ctx := context.Background()

		db, mock, err := sqlmock.New()
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}

		datasql := repository.NewDataSql(db)
		taskRepository := data.NewTaskRepository(datasql.Command())

		us := FindAllTaskUseCaseImpl{
			taskRepository: taskRepository,
		}

		mock.ExpectQuery(
			`SELECT (.+) FROM tb_task t JOIN tb_category c ON NOT c.deleted_at AND t.category_id = c.id WHERE NOT t.deleted_at AND t.user_id = \?`,
		).WithArgs(1).WillReturnRows(sqlmock.NewRows([]string{
			"t.id", "t.created_at", "t.updated_at", "t.uuid", "t.name", "t.description", "t.done", "c.id", "c.created_at", "c.updated_at", "c.slug", "c.name", "t.user_id",
		}).
			AddRow(1, time.Time{}, time.Time{}, "uuid", "task", "new task", false, 1, time.Time{}, time.Time{}, "category", "category", 1).
			AddRow(2, time.Time{}, time.Time{}, "uuid2", "task2", "new task", false, 1, time.Time{}, time.Time{}, "category", "category", 1))

		list, err := us.Execute(ctx, 1)

		if err != nil {
			t.Error(err)
		}

		if len(list) != 2 {
			t.Errorf("list length is %d", len(list))
		}
	})
}
