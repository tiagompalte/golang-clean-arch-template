package usecase

import (
	"context"

	"github.com/tiagompalte/golang-clean-arch-template/pkg/auth"
	"github.com/tiagompalte/golang-clean-arch-template/pkg/errors"
)

type GenerateUserTokenUseCase interface {
	Execute(ctx context.Context, input GenerateUserTokenInput) (GenerateUserTokenOutput, error)
}

type GenerateUserTokenInput struct {
	UUID  string
	Name  string
	Email string
}

type GenerateUserTokenOutput struct {
	AccessToken string
}

type GenerateUserTokenUseCaseImpl struct {
	auth auth.Auth
}

func NewGenerateUserTokenUseCaseImpl(auth auth.Auth) GenerateUserTokenUseCase {
	return GenerateUserTokenUseCaseImpl{
		auth: auth,
	}
}

func (u GenerateUserTokenUseCaseImpl) Execute(ctx context.Context, input GenerateUserTokenInput) (GenerateUserTokenOutput, error) {
	mapClaim := make(map[string]any)
	mapClaim["user_id"] = input.UUID
	mapClaim["user_name"] = input.Name
	mapClaim["user_email"] = input.Email

	token, err := u.auth.GenerateToken(ctx, mapClaim)
	if err != nil {
		return GenerateUserTokenOutput{}, errors.Wrap(err)
	}

	output := GenerateUserTokenOutput{
		AccessToken: token,
	}

	return output, nil
}
