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

// @Summary Find All Tasks
// @Description Find all tasks
// @Tags Task
// @Produce json
// @Security BearerAuth
// @Success 200 {object} []TaskResponse "Tasks list"
// @Router /api/v1/tasks [get]
func FindAllTaskHandler(findAllTaskUseCase usecase.FindAllTaskUseCase) server.Handler {
	return func(w http.ResponseWriter, r *http.Request) error {
		ctx := r.Context()

		user, ok := ctx.Value(constant.ContextUser).(entity.User)
		if !ok {
			return errors.Wrap(pkgErrors.NewInvalidUserError())
		}

		items, err := findAllTaskUseCase.Execute(ctx, user.ID)
		if err != nil {
			return errors.Wrap(err)
		}

		resp := make([]TaskResponse, len(items))
		for i := range items {
			resp[i] = TaskResponse{
				UUID:        items[i].UUID,
				Name:        items[i].Name,
				Description: items[i].Description,
				Done:        items[i].Done,
				Category: CategoryResponse{
					Slug: items[i].Category.GetSlug(),
					Name: items[i].Category.Name,
				},
			}
		}

		err = server.RespondJSON(w, http.StatusOK, resp)
		if err != nil {
			return errors.Wrap(err)
		}

		return nil
	}
}
