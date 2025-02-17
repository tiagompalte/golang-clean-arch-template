package handler

import (
	"net/http"

	"github.com/tiagompalte/golang-clean-arch-template/internal/pkg/server/constant"
	"github.com/tiagompalte/golang-clean-arch-template/internal/pkg/server/middleware"
	"github.com/tiagompalte/golang-clean-arch-template/pkg/errors"
	"github.com/tiagompalte/golang-clean-arch-template/pkg/server"
)

type UserResponse struct {
	UUID    string `json:"id"`
	Version uint32 `json:"version"`
	Name    string `json:"name"`
	Email   string `json:"email"`
}

// @Summary User Logged
// @Description Find user logged
// @Tags User
// @Produce json
// @Security BearerAuth
// @Success 200 {object} UserResponse "User"
// @Router /api/v1/current/user [get]
func FindUserLoggedHandler() server.Handler {
	return func(w http.ResponseWriter, r *http.Request) error {
		ctx := r.Context()

		user, ok := ctx.Value(constant.ContextUser).(middleware.UserToken)
		if !ok {
			return errors.Wrap(errors.NewInvalidUserError())
		}

		resp := UserResponse{
			UUID:    user.UUID,
			Version: user.Version,
			Name:    user.Name,
			Email:   user.Email,
		}

		err := server.RespondJSON(w, http.StatusOK, resp)
		if err != nil {
			return errors.Wrap(err)
		}

		return nil
	}
}
