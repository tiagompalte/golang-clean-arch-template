package handler

import (
	"encoding/json"
	"net/http"

	"github.com/tiagompalte/golang-clean-arch-template/internal/app/usecase"
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

type CreateTaskResponse struct {
	UUID        string                     `json:"uuid"`
	Name        string                     `json:"name"`
	Description string                     `json:"description"`
	Done        bool                       `json:"done"`
	Category    CreateTaskCategoryResponse `json:"category"`
}

type CreateTaskCategoryResponse struct {
	Slug string `json:"slug"`
	Name string `json:"name"`
}

// @Summary Create Task
// @Description Create new Task
// @Tags Task
// @Accept json
// @Produce json
// @Param new_task body CreateTaskRequest true "New Task"
// @Success 201 {object} CreateTaskResponse "Create Task success"
// @Router /api/v1/tasks [post]
func CreateTaskHandler(createTaskUseCase usecase.CreateTask) server.Handler {
	return func(w http.ResponseWriter, r *http.Request) error {
		ctx := r.Context()

		var input CreateTaskRequest
		err := json.NewDecoder(r.Body).Decode(&input)
		if err != nil {
			return errors.Wrap(err)
		}

		task, err := createTaskUseCase.Execute(ctx, input.toInput())
		if err != nil {
			return errors.Wrap(err)
		}

		resp := CreateTaskResponse{
			UUID:        task.UUID,
			Name:        task.Name,
			Description: task.Description,
			Done:        task.Done,
			Category: CreateTaskCategoryResponse{
				Slug: task.Category.GetSlug(),
				Name: task.Category.Name,
			},
		}

		err = server.RespondJSON(w, http.StatusCreated, resp)
		if err != nil {
			return errors.Wrap(err)
		}

		return nil
	}
}
