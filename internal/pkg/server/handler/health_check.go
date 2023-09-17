package handler

import (
	"net/http"

	"github.com/tiagompalte/golang-clean-arch-template/internal/app/usecase"
	"github.com/tiagompalte/golang-clean-arch-template/pkg/errors"
	"github.com/tiagompalte/golang-clean-arch-template/pkg/server"
	usecasePkg "github.com/tiagompalte/golang-clean-arch-template/pkg/usecase"
)

// @Summary Health Check
// @Description Verify health check application
// @Tags Health Check
// @Produce json
// @Success 204
// @Router /api/health-check [get]
func HealthCheckHandler(healthCheck usecase.HealthCheckUseCase) server.Handler {
	return func(w http.ResponseWriter, r *http.Request) error {
		ctx := r.Context()

		_, err := healthCheck.Execute(ctx, usecasePkg.Blank{})
		if err != nil {
			return errors.Wrap(err)
		}

		server.RespondNoContent(w)

		return nil
	}
}
