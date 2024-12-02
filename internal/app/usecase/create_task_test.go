package usecase

import (
	"context"
	"database/sql"
	"testing"
	"time"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"github.com/tiagompalte/golang-clean-arch-template/internal/pkg/infra/uow"
	"github.com/tiagompalte/golang-clean-arch-template/pkg/errors"
	"github.com/tiagompalte/golang-clean-arch-template/pkg/repository"
)

func TestCreateTaskExecute(t *testing.T) {
	t.Parallel()

	t.Run("should return error if is invalid", func(t *testing.T) {
		t.Parallel()
		ctx := context.Background()

		db, _, err := sqlmock.New()
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}

		data := repository.NewDataSql(db)
		uow := uow.NewUow(data)
		us := NewCreateTaskUseCaseImpl(uow)

		input := CreateTaskInput{
			Name:         "",
			Description:  "",
			CategoryName: "",
			UserID:       0,
		}
		_, err = us.Execute(ctx, input)

		if err == nil {
			t.Errorf("should be return error but is nil")
		}
	})

	t.Run("should insert category and task", func(t *testing.T) {
		t.Parallel()
		ctx := context.Background()

		db, mock, err := sqlmock.New()
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}

		data := repository.NewDataSql(db)
		uow := uow.NewUow(data)
		us := NewCreateTaskUseCaseImpl(uow)

		mock.ExpectBegin()

		mock.ExpectQuery(
			`SELECT (.+) FROM tb_category WHERE NOT deleted_at AND slug = \? AND user_id = \?`,
		).WithArgs("category", 1).WillReturnError(sql.ErrNoRows)

		mock.ExpectExec(
			`INSERT INTO tb_category \(slug, name, user_id\) VALUES \(\?,\?,\?\)`,
		).WithArgs("category", "category", 1).WillReturnResult(sqlmock.NewResult(1, 1))

		mock.ExpectExec(
			`INSERT INTO tb_task \(uuid, name, description, category_id, user_id\) VALUES \(\?,\?,\?,\?,\?\)`,
		).WithArgs(sqlmock.AnyArg(), "task", "new task", 1, 1).WillReturnResult(sqlmock.NewResult(1, 1))

		mock.ExpectQuery(
			`SELECT (.+) FROM tb_task t JOIN tb_category c ON NOT c.deleted_at AND t.category_id = c.id WHERE NOT t.deleted_at AND t.id = \?`,
		).WithArgs(1).WillReturnRows(sqlmock.NewRows([]string{
			"t.id", "t.created_at", "t.updated_at", "t.uuid", "t.name", "t.description", "t.done", "c.id", "c.created_at", "c.updated_at", "c.slug", "c.name", "t.user_id",
		}).AddRow(1, time.Time{}, time.Time{}, "uuid", "task", "new task", false, 1, time.Time{}, time.Time{}, "category", "category", 1))

		mock.ExpectCommit()

		input := CreateTaskInput{
			Name:         "task",
			Description:  "new task",
			CategoryName: "category",
			UserID:       uint32(1),
		}
		result, err := us.Execute(ctx, input)

		if err != nil {
			t.Error(err)
		}

		if result.UUID != "uuid" {
			t.Errorf("uuid should be uuid but is %s", result.UUID)
		}

		if result.Done {
			t.Errorf("done should be false but is %t", result.Done)
		}

		if result.CategoryName != "category" {
			t.Errorf("category name should be category but is %s", result.CategoryName)
		}
	})

	t.Run("should insert task and find exists category", func(t *testing.T) {
		t.Parallel()
		ctx := context.Background()

		db, mock, err := sqlmock.New()
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}

		data := repository.NewDataSql(db)
		uow := uow.NewUow(data)
		us := NewCreateTaskUseCaseImpl(uow)

		mock.ExpectBegin()

		mock.ExpectQuery(
			`SELECT (.+) FROM tb_category WHERE NOT deleted_at AND slug = \? AND user_id = \?`,
		).WithArgs("category", 1).WillReturnRows(sqlmock.NewRows([]string{
			"id", "created_at", "updated_at", "slug", "name", "user_id",
		}).AddRow(1, time.Time{}, time.Time{}, "category", "category", 1))

		mock.ExpectExec(
			`INSERT INTO tb_task \(uuid, name, description, category_id, user_id\) VALUES \(\?,\?,\?,\?,\?\)`,
		).WithArgs(sqlmock.AnyArg(), "task", "new task", 1, 1).WillReturnResult(sqlmock.NewResult(1, 1))

		mock.ExpectQuery(
			`SELECT (.+) FROM tb_task t JOIN tb_category c ON NOT c.deleted_at AND t.category_id = c.id WHERE NOT t.deleted_at AND t.id = \?`,
		).WithArgs(1).WillReturnRows(sqlmock.NewRows([]string{
			"t.id", "t.created_at", "t.updated_at", "t.uuid", "t.name", "t.description", "t.done", "c.id", "c.created_at", "c.updated_at", "c.slug", "c.name", "t.user_id",
		}).AddRow(1, time.Time{}, time.Time{}, "uuid", "task", "new task", false, 1, time.Time{}, time.Time{}, "category", "category", 1))

		mock.ExpectCommit()

		input := CreateTaskInput{
			Name:         "task",
			Description:  "new task",
			CategoryName: "category",
			UserID:       uint32(1),
		}
		result, err := us.Execute(ctx, input)

		if err != nil {
			t.Error(err)
		}

		if result.UUID != "uuid" {
			t.Errorf("uuid should be uuid but is %s", result.UUID)
		}

		if result.Done {
			t.Errorf("done should be false but is %t", result.Done)
		}

		if result.CategoryName != "category" {
			t.Errorf("category name should be category but is %s", result.CategoryName)
		}
	})

	t.Run("should rollback if cause error to insert task", func(t *testing.T) {
		t.Parallel()
		ctx := context.Background()

		db, mock, err := sqlmock.New()
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}

		data := repository.NewDataSql(db)
		uow := uow.NewUow(data)
		us := NewCreateTaskUseCaseImpl(uow)

		mock.ExpectBegin()

		mock.ExpectQuery(
			`SELECT (.+) FROM tb_category WHERE NOT deleted_at AND slug = \? AND user_id = \?`,
		).WithArgs("category", 1).WillReturnRows(sqlmock.NewRows([]string{
			"id", "created_at", "updated_at", "slug", "name", "user_id",
		}).AddRow(1, time.Time{}, time.Time{}, "category", "category", 1))

		mock.ExpectExec(
			`INSERT INTO tb_task \(uuid, name, description, category_id, user_id\) VALUES \(\?,\?,\?,\?,\?\)`,
		).WithArgs(sqlmock.AnyArg(), "task", "new task", 1, 1).WillReturnError(errors.NewAppConflictError("task"))

		mock.ExpectRollback()

		input := CreateTaskInput{
			Name:         "task",
			Description:  "new task",
			CategoryName: "category",
			UserID:       uint32(1),
		}
		result, err := us.Execute(ctx, input)

		if !errors.IsAppError(err, errors.ErrorCodeConflict) {
			t.Errorf("error should be conflict but is %s", errors.Cause(err))
		}

		if result.UUID != "" {
			t.Errorf("UUID should be empty but is %s", result.UUID)
		}
	})
}
