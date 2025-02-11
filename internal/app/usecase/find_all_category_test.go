package usecase

import (
	"context"
	"testing"
	"time"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"github.com/tiagompalte/golang-clean-arch-template/internal/pkg/infra/data"
	"github.com/tiagompalte/golang-clean-arch-template/pkg/repository"
)

func TestFindAllCategoryExecute(t *testing.T) {
	t.Parallel()

	t.Run("should return list categories", func(t *testing.T) {
		t.Parallel()
		ctx := context.Background()

		db, mock, err := sqlmock.New()
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}

		datasql := repository.NewDataSql(db)
		categoryRepository := data.NewCategoryRepository(datasql.Command())

		us := FindAllCategoryUseCaseImpl{
			categoryRepository: categoryRepository,
		}

		mock.ExpectQuery(
			`SELECT (.+) FROM tb_category WHERE NOT deleted_at AND user_id = \?`,
		).WithArgs(1).WillReturnRows(sqlmock.NewRows([]string{
			"id", "created_at", "updated_at", "slug", "name", "user_id",
		}).AddRow(1, time.Time{}, time.Time{}, "category", "category", 1).AddRow(2, time.Time{}, time.Time{}, "category2", "category", 1))

		list, err := us.Execute(ctx, 1)

		if err != nil {
			t.Error(err)
		}

		if len(list) != 2 {
			t.Errorf("list length is %d", len(list))
		}
	})
}
