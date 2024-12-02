package handler

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/tiagompalte/golang-clean-arch-template/internal/app/usecase"
	"github.com/tiagompalte/golang-clean-arch-template/internal/pkg/server/constant"
	"github.com/tiagompalte/golang-clean-arch-template/internal/pkg/server/middleware"
	"github.com/tiagompalte/golang-clean-arch-template/pkg/errors"
	"github.com/tiagompalte/golang-clean-arch-template/pkg/server"
)

// @Summary Find One Task
// @Description Find one task by UUID
// @Tags Task
// @Produce json
// @Security BearerAuth
// @Param uuid path string true "Task UUID"
// @Success 200 {object} TaskResponse "Task"
// @Router /api/v1/tasks/{uuid} [get]
func FindOneTaskHandler(findOneTaskUseCase usecase.FindOneTaskUseCase) server.Handler {
	return func(w http.ResponseWriter, r *http.Request) error {
		ctx := r.Context()

		uuid := chi.URLParam(r, "uuid")
		if uuid == "" {
			return errors.NewEmptyPathError("uuid")
		}

		user, ok := ctx.Value(constant.ContextUser).(middleware.UserToken)
		if !ok {
			return errors.Wrap(errors.NewInvalidUserError())
		}

		var input usecase.FindOneTaskInput
		input.TaskUUID = uuid
		input.UserID = user.ID

		task, err := findOneTaskUseCase.Execute(ctx, input)
		if err != nil {
			return errors.Wrap(err)
		}

		resp := TaskResponse{
			UUID:        task.UUID,
			Name:        task.Name,
			Description: task.Description,
			Done:        task.Done,
			Category: CategoryResponse{
				Slug: task.CategorySlug,
				Name: task.CategoryName,
			},
		}

		err = server.RespondJSON(w, http.StatusOK, resp)
		if err != nil {
			return errors.Wrap(err)
		}

		return nil
	}
}
