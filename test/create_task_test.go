//go:build integration
// +build integration

package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tiagompalte/golang-clean-arch-template/internal/pkg/server/constant"
	"github.com/tiagompalte/golang-clean-arch-template/internal/pkg/server/handler"
)

func TestCreateTaskHandler(t *testing.T) {
	t.Parallel()

	t.Run("it should return 201 when created task with success", func(t *testing.T) {
		t.Parallel()

		task := handler.CreateTaskRequest{
			Name:         "Task",
			Description:  "New Task",
			CategoryName: "Category",
		}

		var buf bytes.Buffer
		err := json.NewEncoder(&buf).Encode(task)
		assert.NoError(t, err)

		req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/api/v1/tasks", httpTestUrl), &buf)
		assert.NoError(t, err)
		req.Header.Add(constant.Authorization, fmt.Sprintf("bearer %s", bearerToken))

		resp, err := http.DefaultClient.Do(req)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusCreated, resp.StatusCode)
	})
}
