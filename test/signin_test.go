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
	"github.com/tiagompalte/golang-clean-arch-template/internal/pkg/server/handler"
	"github.com/tiagompalte/golang-clean-arch-template/test/testconfig"
)

func TestSigninHandler(t *testing.T) {
	t.Parallel()

	userLogged, _ := GenerateUserAndToken()
	httpTestUrl := testconfig.Instance().HttpUrl()

	t.Run("it should return 200 and token from user", func(t *testing.T) {
		t.Parallel()

		signinRequest := handler.SigninRequest{
			Email:    userLogged.Email,
			Password: "Pass!1234",
		}

		var buf bytes.Buffer
		err := json.NewEncoder(&buf).Encode(signinRequest)
		assert.NoError(t, err)

		req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/api/v1/signin", httpTestUrl), &buf)
		assert.NoError(t, err)

		resp, err := http.DefaultClient.Do(req)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusCreated, resp.StatusCode)
	})

	t.Run("it should return 401 if informed email and/or password", func(t *testing.T) {
		t.Parallel()

		signinRequest := handler.SigninRequest{
			Email:    "wrong@email.com",
			Password: "wrong_pass123",
		}

		var buf bytes.Buffer
		err := json.NewEncoder(&buf).Encode(signinRequest)
		assert.NoError(t, err)

		req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/api/v1/signin", httpTestUrl), &buf)
		assert.NoError(t, err)

		resp, err := http.DefaultClient.Do(req)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusUnauthorized, resp.StatusCode)
	})
}
