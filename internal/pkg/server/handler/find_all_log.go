package handler

import (
	"net/http"
	"strconv"
	"time"

	"github.com/tiagompalte/golang-clean-arch-template/internal/app/usecase"
	"github.com/tiagompalte/golang-clean-arch-template/pkg/errors"
	"github.com/tiagompalte/golang-clean-arch-template/pkg/server"
)

type LogResponse struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	Level     string    `json:"level"`
	Message   any       `json:"message"`
}

// @Summary Find All Logs
// @Description Find all logs
// @Tags Log
// @Produce json
// @Param limit query int false "Limit" default(100)
// @Success 200 {object} []LogResponse "Logs list"
// @Router /api/v1/logs [get]
func FindAllLogHandler(findAllLogUseCase usecase.FindAllLogUseCase) server.Handler {
	return func(w http.ResponseWriter, r *http.Request) error {
		ctx := r.Context()

		limit := int64(100) // default limit

		if v := r.URL.Query().Get("limit"); v != "" {
			if i, err := strconv.ParseInt(v, 10, 64); err == nil {
				limit = i
			}
		}

		input := usecase.FindAllLogInput{
			Limit: limit,
		}

		logs, err := findAllLogUseCase.Execute(ctx, input)
		if err != nil {
			return errors.Wrap(err)
		}

		resp := make([]LogResponse, len(logs))
		for i := range logs {
			resp[i] = LogResponse{
				ID:        logs[i].ID,
				CreatedAt: logs[i].CreatedAt,
				Level:     logs[i].Level,
				Message:   logs[i].Message,
			}
		}

		err = server.RespondJSON(w, http.StatusOK, resp)
		if err != nil {
			return errors.Wrap(err)
		}

		return nil
	}
}
