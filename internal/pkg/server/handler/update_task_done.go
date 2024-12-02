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

// @Summary Mark task as done
// @Description Update task as done
// @Tags Task
// @Security BearerAuth
// @Param uuid path string true "Task UUID"
// @Success 204
// @Router /api/v1/tasks/{uuid}/done [put]
func UpdateTaskDoneHandler(updateTaskDoneUseCase usecase.UpdateTaskDoneUseCase) server.Handler {
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

		err := updateTaskDoneUseCase.Execute(ctx, usecase.UpdateTaskDoneUseCaseInput{
			UUID:   uuid,
			UserID: user.ID,
		})
		if err != nil {
			return errors.Wrap(err)
		}

		server.RespondNoContent(w)

		return nil
	}
}
