package handler

import (
	"net/http"

	"github.com/tiagompalte/golang-clean-arch-template/internal/app/usecase"
	pkgErrors "github.com/tiagompalte/golang-clean-arch-template/internal/pkg/errors"
	"github.com/tiagompalte/golang-clean-arch-template/pkg/errors"
	"github.com/tiagompalte/golang-clean-arch-template/pkg/server"
)

// @Summary Mark task as undone
// @Description Update task as undone
// @Tags Task
// @Param uuid path string true "Task UUID"
// @Success 204
// @Router /api/v1/tasks/{uuid}/undone [put]
func UpdateTaskUndoneHandler(updateTaskUndoneUseCase usecase.UpdateTaskUndone) server.Handler {
	return func(w http.ResponseWriter, r *http.Request) error {
		ctx := r.Context()

		uuid, ok := extractParamPath(r, 4)
		if !ok {
			return pkgErrors.NewEmptyPathError("uuid")
		}

		_, err := updateTaskUndoneUseCase.Execute(ctx, uuid)
		if err != nil {
			return errors.Wrap(err)
		}

		server.RespondNoContent(w)

		return nil
	}
}
