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

type UpdateUserName struct {
	Version uint32 `json:"version"`
	NewName string `json:"new_name"`
}

// @Summary     Update User Name
// @Description Update user name
// @Tags        User
// @Accept      json
// @Produce     json
// @Security    BearerAuth
// @Param       update_user_name body     UpdateUserName true "Update User Name"
// @Success     200              {object} UserResponse   "User"
// @Router      /api/v1/current/user/update-name [put]
func UpdateUserNameHandler(updateUserNameUseCase usecase.UpdateUserNameUseCase) server.Handler {
	return func(w http.ResponseWriter, r *http.Request) error {
		ctx := r.Context()

		userToken, ok := ctx.Value(constant.ContextUser).(middleware.UserToken)
		if !ok {
			return errors.Wrap(errors.NewInvalidUserError())
		}

		var inputBody UpdateUserName
		err := json.NewDecoder(r.Body).Decode(&inputBody)
		if err != nil {
			return errors.Wrap(err)
		}

		input := usecase.UpdateUserNameInput{
			UserUUID:    userToken.UUID,
			Version:     inputBody.Version,
			NewUserName: inputBody.NewName,
		}

		user, err := updateUserNameUseCase.Execute(ctx, input)
		if err != nil {
			return errors.Wrap(err)
		}

		resp := UserResponse{
			UUID:    user.UUID,
			Version: user.Version,
			Name:    user.Name,
			Email:   user.Email,
		}

		err = server.RespondJSON(w, http.StatusOK, resp)
		if err != nil {
			return errors.Wrap(err)
		}

		return nil
	}
}
