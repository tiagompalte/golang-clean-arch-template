package usecase

import (
	"context"

	"github.com/tiagompalte/golang-clean-arch-template/internal/app/entity"
	"github.com/tiagompalte/golang-clean-arch-template/pkg/auth"
	"github.com/tiagompalte/golang-clean-arch-template/pkg/errors"
)

type GenerateUserTokenUseCase interface {
	Execute(ctx context.Context, user entity.User) (GenerateUserTokenOutput, error)
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

func (u GenerateUserTokenUseCaseImpl) Execute(ctx context.Context, user entity.User) (GenerateUserTokenOutput, error) {
	mapClaim := make(map[string]any)
	mapClaim["user_id"] = user.UUID
	mapClaim["user_name"] = user.Name
	mapClaim["user_email"] = user.Email

	token, err := u.auth.GenerateToken(ctx, mapClaim)
	if err != nil {
		return GenerateUserTokenOutput{}, errors.Wrap(err)
	}

	output := GenerateUserTokenOutput{
		AccessToken: token,
	}

	return output, nil
}
