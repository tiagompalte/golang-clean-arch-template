//go:build integration
// +build integration

package test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tiagompalte/golang-clean-arch-template/test/testconfig"
)

func TestHealthCheck(t *testing.T) {
	t.Parallel()

	httpTestUrl := testconfig.Instance().HttpUrl()

	t.Run("it should return 204 when health is ok", func(t *testing.T) {
		t.Parallel()

		resp, err := http.Get(fmt.Sprintf("%s/api/health-check", httpTestUrl))

		assert.NoError(t, err)
		assert.Equal(t, http.StatusNoContent, resp.StatusCode)
	})
}
