package handler

import (
	"net/http"
	"time"

	"github.com/tiagompalte/golang-clean-arch-template/internal/app/entity"
	pkgErrors "github.com/tiagompalte/golang-clean-arch-template/internal/pkg/errors"
	"github.com/tiagompalte/golang-clean-arch-template/internal/pkg/server/constant"
	"github.com/tiagompalte/golang-clean-arch-template/pkg/errors"
	"github.com/tiagompalte/golang-clean-arch-template/pkg/server"
)

type UserResponse struct {
	UUID      string    `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
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

		user, ok := ctx.Value(constant.ContextUser).(entity.User)
		if !ok {
			return errors.Wrap(pkgErrors.NewInvalidUserError())
		}

		resp := UserResponse{
			UUID:      user.UUID,
			CreatedAt: user.CreatedAt,
			Name:      user.Name,
			Email:     user.Email,
		}

		err := server.RespondJSON(w, http.StatusOK, resp)
		if err != nil {
			return errors.Wrap(err)
		}

		return nil
	}
}
