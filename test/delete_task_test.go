//go:build integration
// +build integration

package test

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tiagompalte/golang-clean-arch-template/internal/app/usecase"
	"github.com/tiagompalte/golang-clean-arch-template/internal/pkg/server/constant"
	"github.com/tiagompalte/golang-clean-arch-template/test/testconfig"
)

func TestDeleteTaskHandler(t *testing.T) {
	t.Parallel()

	userLogged, token := GenerateUserAndToken()
	httpTestUrl := testconfig.Instance().HttpUrl()
	app := testconfig.Instance().App()

	t.Run("it should return 204 when deleted task with success", func(t *testing.T) {
		t.Parallel()
		ctx := context.Background()

		createTask := usecase.CreateTaskInput{
			Name:         "Task deleted",
			Description:  "Task deleted",
			CategoryName: "Deleted",
			UserID:       userLogged.ID,
		}

		task, err := app.UseCase().CreateTaskUseCase.Execute(ctx, createTask)
		assert.NoError(t, err)

		req, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("%s/api/v1/tasks/%s", httpTestUrl, task.UUID), nil)
		assert.NoError(t, err)
		req.Header.Add(constant.Authorization, fmt.Sprintf("bearer %s", token))

		resp, err := http.DefaultClient.Do(req)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusNoContent, resp.StatusCode)
	})

	t.Run("it should return 401 when informed token from another user", func(t *testing.T) {
		t.Parallel()

		anotherUser, _ := GenerateUserAndToken()

		ctx := context.Background()

		createTask := usecase.CreateTaskInput{
			Name:         "Task deleted",
			Description:  "Task deleted",
			CategoryName: "Deleted",
			UserID:       anotherUser.ID,
		}

		task, err := app.UseCase().CreateTaskUseCase.Execute(ctx, createTask)
		assert.NoError(t, err)

		req, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("%s/api/v1/tasks/%s", httpTestUrl, task.UUID), nil)
		assert.NoError(t, err)
		req.Header.Add(constant.Authorization, fmt.Sprintf("bearer %s", token))

		resp, err := http.DefaultClient.Do(req)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusUnauthorized, resp.StatusCode)
	})
}
