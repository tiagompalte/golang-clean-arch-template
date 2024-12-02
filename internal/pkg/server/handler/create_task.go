package handler

import (
	"encoding/json"
	"net/http"

	"github.com/tiagompalte/golang-clean-arch-template/internal/app/usecase"
	"github.com/tiagompalte/golang-clean-arch-template/internal/pkg/server/constant"
	"github.com/tiagompalte/golang-clean-arch-template/internal/pkg/server/middleware"
	"github.com/tiagompalte/golang-clean-arch-template/pkg/errors"
	"github.com/tiagompalte/golang-clean-arch-template/pkg/server"
)

type CreateTaskRequest struct {
	Name         string `json:"name"`
	Description  string `json:"description"`
	CategoryName string `json:"category_name"`
}

func (r CreateTaskRequest) toInput() usecase.CreateTaskInput {
	return usecase.CreateTaskInput{
		Name:         r.Name,
		Description:  r.Description,
		CategoryName: r.CategoryName,
	}
}

type TaskResponse struct {
	UUID        string           `json:"uuid"`
	Name        string           `json:"name"`
	Description string           `json:"description"`
	Done        bool             `json:"done"`
	Category    CategoryResponse `json:"category"`
}

// @Summary Create Task
// @Description Create new Task
// @Tags Task
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param new_task body CreateTaskRequest true "New Task"
// @Success 201 {object} TaskResponse "Create Task success"
// @Router /api/v1/tasks [post]
func CreateTaskHandler(createTaskUseCase usecase.CreateTaskUseCase) server.Handler {
	return func(w http.ResponseWriter, r *http.Request) error {
		ctx := r.Context()

		var request CreateTaskRequest
		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			return errors.Wrap(err)
		}

		user, ok := ctx.Value(constant.ContextUser).(middleware.UserToken)
		if !ok {
			return errors.Wrap(errors.NewInvalidUserError())
		}

		input := request.toInput()
		input.UserID = user.ID

		task, err := createTaskUseCase.Execute(ctx, input)
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

		err = server.RespondJSON(w, http.StatusCreated, resp)
		if err != nil {
			return errors.Wrap(err)
		}

		return nil
	}
}
