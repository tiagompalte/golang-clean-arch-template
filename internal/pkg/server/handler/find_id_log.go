package handler

import (
	"net/http"

	"github.com/tiagompalte/golang-clean-arch-template/internal/app/usecase"
	"github.com/tiagompalte/golang-clean-arch-template/pkg/errors"
	"github.com/tiagompalte/golang-clean-arch-template/pkg/server"
)

// @Summary Find Log by ID
// @Description Find log by ID
// @Tags Log
// @Produce json
// @Param id path string true "Log ID"
// @Success 200 {object} LogResponse "Log found"
// @Router /api/v1/logs/{id} [get]
func FindByIDLogHandler(findByIDLogUseCase usecase.FindByIDLogUseCase) server.Handler {
	return func(w http.ResponseWriter, r *http.Request) error {
		ctx := r.Context()

		id := r.PathValue("id")

		log, err := findByIDLogUseCase.Execute(ctx, id)
		if err != nil {
			return errors.Wrap(err)
		}

		var response LogResponse
		response.ID = log.ID.Hex()
		response.CreatedAt = log.CreatedAt
		response.Level = log.Level
		response.Message = log.Message

		err = server.RespondJSON(w, http.StatusOK, response)
		if err != nil {
			return errors.Wrap(err)
		}

		return nil
	}
}
