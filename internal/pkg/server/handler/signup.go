package handler

import (
	"encoding/json"
	"net/http"

	"github.com/tiagompalte/golang-clean-arch-template/internal/app/usecase"
	"github.com/tiagompalte/golang-clean-arch-template/pkg/errors"
	"github.com/tiagompalte/golang-clean-arch-template/pkg/server"
)

type SignupRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SignupResponse struct {
	AccessToken string `json:"access_token"`
}

func (r SignupRequest) toInput() usecase.CreateUserInput {
	return usecase.CreateUserInput{
		Name:     r.Name,
		Email:    r.Email,
		Password: r.Password,
	}
}

// @Summary Sign Up
// @Description Create new User
// @Tags Sign Up
// @Accept json
// @Produce json
// @Param signup body SignupRequest true "New User"
// @Success 201 {object} SignupResponse "Tokens"
// @Router /api/v1/signup [post]
func SignupHandler(
	createUserUseCase usecase.CreateUserUseCase,
	generateUserTokenUseCase usecase.GenerateUserTokenUseCase,
) server.Handler {
	return func(w http.ResponseWriter, r *http.Request) error {
		ctx := r.Context()

		var input SignupRequest
		err := json.NewDecoder(r.Body).Decode(&input)
		if err != nil {
			return errors.Wrap(err)
		}

		user, err := createUserUseCase.Execute(ctx, input.toInput())
		if err != nil {
			return errors.Wrap(err)
		}

		output, err := generateUserTokenUseCase.Execute(ctx, user)
		if err != nil {
			return errors.Wrap(err)
		}

		response := SignupResponse{
			AccessToken: output.AccessToken,
		}

		err = server.RespondJSON(w, http.StatusCreated, response)
		if err != nil {
			return errors.Wrap(err)
		}

		return nil
	}
}
