package usecase

import (
	"context"

	"github.com/tiagompalte/golang-clean-arch-template/internal/app/entity"
	"github.com/tiagompalte/golang-clean-arch-template/internal/app/repository"
	errPkg "github.com/tiagompalte/golang-clean-arch-template/internal/pkg/errors"
	"github.com/tiagompalte/golang-clean-arch-template/pkg/crypto"
	"github.com/tiagompalte/golang-clean-arch-template/pkg/errors"
)

type CreateUserUseCase interface {
	Execute(ctx context.Context, input CreateUserInput) (entity.User, error)
}

type CreateUserInput struct {
	Name     string
	Email    string
	Password string
}

func (i CreateUserInput) Validate() error {
	aggrErr := errors.NewAggregatedError()
	if i.Name == "" {
		aggrErr.Add(errPkg.NewEmptyParameterError("name"))
	}
	if i.Email == "" {
		aggrErr.Add(errPkg.NewEmptyParameterError("email"))
	}
	if i.Password == "" {
		aggrErr.Add(errPkg.NewEmptyParameterError("password"))
	}

	if aggrErr.Len() > 0 {
		return errors.Wrap(aggrErr)
	}

	return nil
}

type CreateUserUseCaseImpl struct {
	userRepository repository.UserRepository
	crypto         crypto.Crypto
}

func NewCreateUserUseCaseImpl(userRepository repository.UserRepository, crypto crypto.Crypto) CreateUserUseCase {
	return CreateUserUseCaseImpl{
		userRepository: userRepository,
		crypto:         crypto,
	}
}

func (u CreateUserUseCaseImpl) Execute(ctx context.Context, input CreateUserInput) (entity.User, error) {
	err := input.Validate()
	if err != nil {
		return entity.User{}, errors.Wrap(err)
	}

	var user entity.User
	user.Name = input.Name
	user.Email = input.Email

	passEncrypted, err := u.crypto.GenerateHash(ctx, input.Password)
	if err != nil {
		return entity.User{}, errors.Wrap(err)
	}

	id, err := u.userRepository.Insert(ctx, user, passEncrypted)
	if err != nil {
		return entity.User{}, errors.Wrap(err)
	}

	user, err = u.userRepository.FindByID(ctx, id)
	if err != nil {
		return entity.User{}, errors.Wrap(err)
	}

	return user, nil
}
