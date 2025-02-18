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

func TestFindOneTaskHandler(t *testing.T) {
	t.Parallel()

	user, token := GenerateUserAndToken()
	httpTestUrl := testconfig.Instance().HttpUrl()
	app := testconfig.Instance().App()

	ctx := context.Background()

	task, err := app.UseCase().CreateTaskUseCase.Execute(ctx, usecase.CreateTaskInput{
		Name:         "Task Only Item",
		Description:  "Description",
		CategoryName: "Category Task Only",
		UserID:       user.ID,
	})
	assert.NoError(t, err)

	t.Run("it should return 200 and return task by uuid", func(t *testing.T) {
		t.Parallel()

		req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/api/v1/tasks/%s", httpTestUrl, task.UUID), nil)
		assert.NoError(t, err)
		req.Header.Add(constant.Authorization, fmt.Sprintf("bearer %s", token))

		resp, err := http.DefaultClient.Do(req)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, resp.StatusCode)

		var respBody handler.TaskResponse
		err = json.NewDecoder(resp.Body).Decode(&respBody)
		assert.NoError(t, err)

		assert.Equal(t, task.UUID, respBody.UUID)
		assert.Equal(t, task.Name, respBody.Name)
		assert.Equal(t, task.Description, respBody.Description)
		assert.Equal(t, task.Done, respBody.Done)
		assert.Equal(t, task.CategoryName, respBody.Category.Name)
		assert.Equal(t, task.CategorySlug, respBody.Category.Slug)

	})

}
