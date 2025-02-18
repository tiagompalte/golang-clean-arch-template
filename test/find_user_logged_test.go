//go:build integration
// +build integration

package test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tiagompalte/golang-clean-arch-template/internal/pkg/server/constant"
	"github.com/tiagompalte/golang-clean-arch-template/internal/pkg/server/handler"
	"github.com/tiagompalte/golang-clean-arch-template/test/testconfig"
)

func TestFindUserLoggedHandler(t *testing.T) {
	t.Parallel()

	httpTestUrl := testconfig.Instance().HttpUrl()

	t.Run("it should return 200 and return user logged", func(t *testing.T) {
		t.Parallel()

		user, token := GenerateUserAndToken()

		req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/api/v1/current/user", httpTestUrl), nil)
		assert.NoError(t, err)
		req.Header.Add(constant.Authorization, fmt.Sprintf("bearer %s", token))

		resp, err := http.DefaultClient.Do(req)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, resp.StatusCode)

		var respBody handler.UserResponse
		err = json.NewDecoder(resp.Body).Decode(&respBody)
		assert.NoError(t, err)

		assert.Equal(t, user.UUID, respBody.UUID)
		assert.Equal(t, user.Name, respBody.Name)
		assert.Equal(t, user.Email, respBody.Email)
	})

}
