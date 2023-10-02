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

		result, err := findAllTaskUseCase.Execute(ctx, user.ID)
		if err != nil {
			return errors.Wrap(err)
		}

		resp := make([]TaskResponse, len(result.Items))
		for i := range result.Items {
			resp[i] = TaskResponse{
				UUID:        result.Items[i].UUID,
				Name:        result.Items[i].Name,
				Description: result.Items[i].Description,
				Done:        result.Items[i].Done,
				Category: CategoryResponse{
					Slug: result.Items[i].Category.GetSlug(),
					Name: result.Items[i].Category.Name,
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
