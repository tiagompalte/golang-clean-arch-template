package handler

import (
	"net/http"

	"github.com/tiagompalte/golang-clean-arch-template/internal/app/usecase"
	pkgErrors "github.com/tiagompalte/golang-clean-arch-template/internal/pkg/errors"
	"github.com/tiagompalte/golang-clean-arch-template/pkg/errors"
	"github.com/tiagompalte/golang-clean-arch-template/pkg/server"
)

// @Summary Mark task as done
// @Description Update task as done
// @Tags Task
// @Param uuid path string true "Task UUID"
// @Success 204
// @Router /api/v1/tasks/{uuid}/done [put]
func UpdateTaskDoneHandler(updateTaskDoneUseCase usecase.UpdateTaskDone) server.Handler {
	return func(w http.ResponseWriter, r *http.Request) error {
		ctx := r.Context()

		uuid, ok := extractParamPath(r, 4)
		if !ok {
			return pkgErrors.NewEmptyPathError("uuid")
		}

		_, err := updateTaskDoneUseCase.Execute(ctx, uuid)
		if err != nil {
			return errors.Wrap(err)
		}

		server.RespondNoContent(w)

		return nil
	}
}
