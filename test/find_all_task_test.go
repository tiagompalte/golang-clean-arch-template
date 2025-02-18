//go:build integration
// +build integration

package test

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tiagompalte/golang-clean-arch-template/internal/app/usecase"
	"github.com/tiagompalte/golang-clean-arch-template/internal/pkg/server/constant"
	"github.com/tiagompalte/golang-clean-arch-template/internal/pkg/server/handler"
	"github.com/tiagompalte/golang-clean-arch-template/test/testconfig"
)

func TestFindAllTaskHandler(t *testing.T) {
	t.Parallel()

	user, token := GenerateUserAndToken()
	httpTestUrl := testconfig.Instance().HttpUrl()
	app := testconfig.Instance().App()

	ctx := context.Background()

	task1, err := app.UseCase().CreateTaskUseCase.Execute(ctx, usecase.CreateTaskInput{
		Name:         "Task Item List 1",
		Description:  "Description",
		CategoryName: "Category Task 1",
		UserID:       user.ID,
	})
	assert.NoError(t, err)

	task2, err := app.UseCase().CreateTaskUseCase.Execute(ctx, usecase.CreateTaskInput{
		Name:         "Task Item List 2",
		Description:  "Description",
		CategoryName: "Category Task 2",
		UserID:       user.ID,
	})
	assert.NoError(t, err)

	t.Run("it should return 200 and return tasks list", func(t *testing.T) {
		t.Parallel()

		req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/api/v1/tasks", httpTestUrl), nil)
		assert.NoError(t, err)
		req.Header.Add(constant.Authorization, fmt.Sprintf("bearer %s", token))

		resp, err := http.DefaultClient.Do(req)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, resp.StatusCode)

		var respBody []handler.TaskResponse
		err = json.NewDecoder(resp.Body).Decode(&respBody)
		assert.NoError(t, err)

		for _, item := range respBody {
			assert.True(t, item.UUID == task1.UUID || item.UUID == task2.UUID)
			assert.True(t, item.Name == task1.Name || item.Name == task2.Name)
			assert.True(t, item.Description == task1.Description || item.Description == task2.Description)
			assert.True(t, item.Done == task1.Done || item.Done == task2.Done)
			assert.True(t, item.Category.Name == task1.CategoryName || item.Category.Name == task2.CategoryName)
			assert.True(t, item.Category.Slug == task1.CategorySlug || item.Category.Slug == task2.CategorySlug)
		}
	})
}
