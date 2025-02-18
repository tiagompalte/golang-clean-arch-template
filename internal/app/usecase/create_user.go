package usecase

import (
	"context"

	"github.com/tiagompalte/golang-clean-arch-template/internal/app/entity"
	"github.com/tiagompalte/golang-clean-arch-template/internal/app/protocols"
	"github.com/tiagompalte/golang-clean-arch-template/pkg/crypto"
	"github.com/tiagompalte/golang-clean-arch-template/pkg/errors"
)

type CreateUserUseCase interface {
	Execute(ctx context.Context, input CreateUserInput) (CreateUserOutput, error)
}

type CreateUserInput struct {
	Name     string
	Email    string
	Password string
}

type CreateUserOutput struct {
	ID      uint32
	Version uint32
	UUID    string
	Name    string
	Email   string
}

type CreateUserUseCaseImpl struct {
	userRepository protocols.UserRepository
	crypto         crypto.Crypto
}

func NewCreateUserUseCaseImpl(userRepository protocols.UserRepository, crypto crypto.Crypto) CreateUserUseCase {
	return CreateUserUseCaseImpl{
		userRepository: userRepository,
		crypto:         crypto,
	}
}

func (u CreateUserUseCaseImpl) Execute(ctx context.Context, input CreateUserInput) (CreateUserOutput, error) {
	var user entity.User
	user.Name = input.Name
	user.Email = input.Email

	err := user.ValidateNew()
	if err != nil {
		return CreateUserOutput{}, errors.Wrap(err)
	}

	if input.Password == "" {
		return CreateUserOutput{}, errors.Wrap(errors.NewEmptyParameterError("password"))
	}

	passEncrypted, err := u.crypto.GenerateHash(ctx, input.Password)
	if err != nil {
		return CreateUserOutput{}, errors.Wrap(err)
	}

	id, err := u.userRepository.Insert(ctx, user, passEncrypted)
	if err != nil {
		return CreateUserOutput{}, errors.Wrap(err)
	}

	user, err = u.userRepository.FindByID(ctx, id)
	if err != nil {
		return CreateUserOutput{}, errors.Wrap(err)
	}

	return CreateUserOutput{
		ID:      user.ID,
		Version: user.Version,
		UUID:    user.UUID,
		Name:    user.Name,
		Email:   user.Email,
	}, nil
}
