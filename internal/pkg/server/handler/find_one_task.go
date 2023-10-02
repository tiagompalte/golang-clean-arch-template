package handler

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/tiagompalte/golang-clean-arch-template/internal/app/entity"
	"github.com/tiagompalte/golang-clean-arch-template/internal/app/usecase"
	pkgErrors "github.com/tiagompalte/golang-clean-arch-template/internal/pkg/errors"
	"github.com/tiagompalte/golang-clean-arch-template/internal/pkg/server/constant"
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
			return pkgErrors.NewEmptyPathError("uuid")
		}

		user, ok := ctx.Value(constant.ContextUser).(entity.User)
		if !ok {
			return errors.Wrap(pkgErrors.NewInvalidUserError())
		}

		task, err := findOneTaskUseCase.Execute(ctx, uuid)
		if err != nil {
			return errors.Wrap(err)
		}

		if task.UserID != user.ID {
			return errors.Wrap(pkgErrors.NewInvalidUserError())
		}

		resp := TaskResponse{
			UUID:        task.UUID,
			Name:        task.Name,
			Description: task.Description,
			Done:        task.Done,
			Category: CategoryResponse{
				Slug: task.Category.GetSlug(),
				Name: task.Category.Name,
			},
		}

		err = server.RespondJSON(w, http.StatusOK, resp)
		if err != nil {
			return errors.Wrap(err)
		}

		return nil
	}
}
