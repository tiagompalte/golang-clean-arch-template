package handler

import (
	"net/http"

	"github.com/tiagompalte/golang-clean-arch-template/internal/app/usecase"
	"github.com/tiagompalte/golang-clean-arch-template/pkg/errors"
	"github.com/tiagompalte/golang-clean-arch-template/pkg/server"
)

// @Summary Find All Tasks
// @Description Find all tasks
// @Tags Task
// @Produce json
// @Success 200 {object} []TaskResponse "Tasks list"
// @Router /api/v1/tasks [get]
func FindAllTaskHandler(findAllTaskUseCase usecase.FindAllTaskUseCase) server.Handler {
	return func(w http.ResponseWriter, r *http.Request) error {
		ctx := r.Context()

		result, err := findAllTaskUseCase.Execute(ctx, blank)
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
