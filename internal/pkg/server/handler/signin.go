package handler

import (
	"encoding/json"
	"net/http"

	"github.com/tiagompalte/golang-clean-arch-template/internal/app/usecase"
	"github.com/tiagompalte/golang-clean-arch-template/pkg/errors"
	"github.com/tiagompalte/golang-clean-arch-template/pkg/server"
)

type SigninRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (r SigninRequest) toInput() usecase.ValidateUserPasswordInput {
	return usecase.ValidateUserPasswordInput{
		Email:    r.Email,
		Password: r.Password,
	}
}

type SigninResponse struct {
	AccessToken string `json:"access_token"`
}

// @Summary Sign In
// @Description Login user
// @Tags Sign In
// @Accept json
// @Produce json
// @Param signin body SigninRequest true "Login User"
// @Success 201 {object} SigninResponse "Tokens"
// @Router /api/v1/signin [post]
func SigninHandler(
	validateUserPasswordUseCase usecase.ValidateUserPasswordUseCase,
	generateUserTokenUseCase usecase.GenerateUserTokenUseCase,
) server.Handler {
	return func(w http.ResponseWriter, r *http.Request) error {
		ctx := r.Context()

		var signinRequest SigninRequest
		err := json.NewDecoder(r.Body).Decode(&signinRequest)
		if err != nil {
			return errors.Wrap(err)
		}

		user, err := validateUserPasswordUseCase.Execute(ctx, signinRequest.toInput())
		if err != nil {
			return errors.Wrap(err)
		}

		var input usecase.GenerateUserTokenInput
		input.UUID = user.UUID
		input.Name = user.Name
		input.Email = user.Email

		output, err := generateUserTokenUseCase.Execute(ctx, input)
		if err != nil {
			return errors.Wrap(err)
		}

		response := SigninResponse{
			AccessToken: output.AccessToken,
		}

		err = server.RespondJSON(w, http.StatusCreated, response)
		if err != nil {
			return errors.Wrap(err)
		}

		return nil
	}
}
