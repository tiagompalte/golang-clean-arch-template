package handler

import (
	"net/http"

	"github.com/tiagompalte/golang-clean-arch-template/internal/app/entity"
	"github.com/tiagompalte/golang-clean-arch-template/internal/app/usecase"
	pkgErrors "github.com/tiagompalte/golang-clean-arch-template/internal/pkg/errors"
	"github.com/tiagompalte/golang-clean-arch-template/internal/pkg/server/constant"
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
// @Security BearerAuth
// @Success 200 {object} []CategoryResponse "Categories list"
// @Router /api/v1/categories [get]
func FindAllCategoryHandler(findAllCategoryUseCase usecase.FindAllCategoryUseCase) server.Handler {
	return func(w http.ResponseWriter, r *http.Request) error {
		ctx := r.Context()

		user, ok := ctx.Value(constant.ContextUser).(entity.User)
		if !ok {
			return errors.Wrap(pkgErrors.NewInvalidUserError())
		}

		items, err := findAllCategoryUseCase.Execute(ctx, user.ID)
		if err != nil {
			return errors.Wrap(err)
		}

		resp := make([]CategoryResponse, len(items))
		for i := range items {
			resp[i] = CategoryResponse{
				Slug: items[i].GetSlug(),
				Name: items[i].Name,
			}
		}

		err = server.RespondJSON(w, http.StatusOK, resp)
		if err != nil {
			return errors.Wrap(err)
		}

		return nil
	}
}
