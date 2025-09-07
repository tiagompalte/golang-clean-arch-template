package handler

import (
	"encoding/json"
	"net/http"

	"github.com/tiagompalte/golang-clean-arch-template/internal/app/usecase"
	"github.com/tiagompalte/golang-clean-arch-template/pkg/errors"
	"github.com/tiagompalte/golang-clean-arch-template/pkg/server"
)

type CreateLogRequest struct {
	Level   string `json:"level"`
	Message string `json:"message"`
}

type LogResponse struct {
	ID any `json:"id"`
}

func (r *CreateLogRequest) toInput() usecase.CreateLogInput {
	return usecase.CreateLogInput{
		Level:   r.Level,
		Message: r.Message,
	}
}

// @Summary Create Log
// @Description Create new Log
// @Tags Log
// @Accept json
// @Produce json
// @Param new_log body CreateLogRequest true "New Log"
// @Success 201 {object} LogResponse "Create Log success"
// @Router /api/v1/logs [post]
func CreateLogHandler(createLogUseCase usecase.CreateLogUseCase) server.Handler {
	return func(w http.ResponseWriter, r *http.Request) error {
		ctx := r.Context()

		var request CreateLogRequest
		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			return errors.Wrap(err)
		}

		input := request.toInput()

		output, err := createLogUseCase.Execute(ctx, input)
		if err != nil {
			return errors.Wrap(err)
		}

		response := LogResponse{
			ID: output.ID,
		}

		err = server.RespondJSON(w, http.StatusCreated, response)
		if err != nil {
			return errors.Wrap(err)
		}

		return nil
	}
}
