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

func TestFindAllCategoryHandler(t *testing.T) {
	t.Parallel()

	user, token := GenerateUserAndToken()
	httpTestUrl := testconfig.Instance().HttpUrl()
	app := testconfig.Instance().App()

	ctx := context.Background()

	category1, err := app.UseCase().CreateCategoryUseCase.Execute(ctx, usecase.CreateCategoryInput{
		Name:   "Category Item List 1",
		UserID: user.ID,
	})
	assert.NoError(t, err)

	category2, err := app.UseCase().CreateCategoryUseCase.Execute(ctx, usecase.CreateCategoryInput{
		Name:   "Category Item List 2",
		UserID: user.ID,
	})
	assert.NoError(t, err)

	t.Run("it should return 200 and return categories list", func(t *testing.T) {
		t.Parallel()

		req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/api/v1/categories", httpTestUrl), nil)
		assert.NoError(t, err)
		req.Header.Add(constant.Authorization, fmt.Sprintf("bearer %s", token))

		resp, err := http.DefaultClient.Do(req)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, resp.StatusCode)

		var respBody []handler.CategoryResponse
		err = json.NewDecoder(resp.Body).Decode(&respBody)
		assert.NoError(t, err)

		for _, item := range respBody {
			assert.True(t, item.Name == category1.Name || item.Name == category2.Name)
			assert.True(t, item.Slug == category1.Slug || item.Slug == category2.Slug)
		}
	})
}
