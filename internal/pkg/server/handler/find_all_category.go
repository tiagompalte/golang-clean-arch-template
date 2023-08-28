package handler

import (
	"net/http"

	"github.com/tiagompalte/golang-clean-arch-template/internal/app/usecase"
	"github.com/tiagompalte/golang-clean-arch-template/pkg/errors"
	"github.com/tiagompalte/golang-clean-arch-template/pkg/server"
)

type CategoryResponse struct {
	Slug string `json:"slug"`
	Name string `json:"name"`
}

// @Summary Find All Categories
// @Description Find all categories
// @Tags Category
// @Produce json
// @Success 200 {object} []CategoryResponse "Categories list"
// @Router /api/v1/categories [get]
func FindAllCategoryHandler(findAllCategoryUseCase usecase.FindAllCategory) server.Handler {
	return func(w http.ResponseWriter, r *http.Request) error {
		ctx := r.Context()

		result, err := findAllCategoryUseCase.Execute(ctx, blank)
		if err != nil {
			return errors.Wrap(err)
		}

		resp := make([]CategoryResponse, len(result.Items))
		for i := range result.Items {
			resp[i] = CategoryResponse{
				Slug: result.Items[i].GetSlug(),
				Name: result.Items[i].Name,
			}
		}

		err = server.RespondJSON(w, http.StatusOK, resp)
		if err != nil {
			return errors.Wrap(err)
		}

		return nil
	}
}
