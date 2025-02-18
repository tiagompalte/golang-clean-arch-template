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
	"github.com/tiagompalte/golang-clean-arch-template/pkg/errors"
	"github.com/tiagompalte/golang-clean-arch-template/pkg/server"
	"github.com/tiagompalte/golang-clean-arch-template/test/testconfig"
)

func TestUpdateUserNameHandler(t *testing.T) {
	t.Parallel()

	httpTestUrl := testconfig.Instance().HttpUrl()

	t.Run("it should update user name", func(t *testing.T) {
		t.Parallel()

		user, token := GenerateUserAndToken()

		updateUserName := handler.UpdateUserName{
			Version: user.Version,
			NewName: "New User Name",
		}

		var buf bytes.Buffer
		err := json.NewEncoder(&buf).Encode(updateUserName)
		assert.NoError(t, err)

		req, err := http.NewRequest(http.MethodPut, fmt.Sprintf("%s/api/v1/current/user/update-name", httpTestUrl), &buf)
		assert.NoError(t, err)
		req.Header.Add(constant.Authorization, fmt.Sprintf("bearer %s", token))

		resp, err := http.DefaultClient.Do(req)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, resp.StatusCode)

		var respBody handler.UserResponse
		err = json.NewDecoder(resp.Body).Decode(&respBody)
		assert.NoError(t, err)

		assert.Equal(t, user.UUID, respBody.UUID)
		assert.Equal(t, updateUserName.NewName, respBody.Name)
		assert.Equal(t, user.Email, respBody.Email)
		assert.Equal(t, updateUserName.Version+1, respBody.Version)
	})

	t.Run("it should return 400 when version is invalid", func(t *testing.T) {
		t.Parallel()

		user, token := GenerateUserAndToken()

		updateUserName := handler.UpdateUserName{
			Version: user.Version + 1,
			NewName: "New User Name",
		}

		var buf bytes.Buffer
		err := json.NewEncoder(&buf).Encode(updateUserName)
		assert.NoError(t, err)

		req, err := http.NewRequest(http.MethodPut, fmt.Sprintf("%s/api/v1/current/user/update-name", httpTestUrl), &buf)
		assert.NoError(t, err)
		req.Header.Add(constant.Authorization, fmt.Sprintf("bearer %s", token))

		resp, err := http.DefaultClient.Do(req)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusBadRequest, resp.StatusCode)

		var errorResponse server.ErrorResponseWrapper
		err = json.NewDecoder(resp.Body).Decode(&errorResponse)
		assert.NoError(t, err)

		assert.Equal(t, errors.ErrorCodeConcurrencyRepository, errorResponse.Errors[0].Code)
	})
}
