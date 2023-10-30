package usecase

import (
	"context"

	"github.com/tiagompalte/golang-clean-arch-template/internal/app/entity"
	"github.com/tiagompalte/golang-clean-arch-template/internal/app/repository"
	errPkg "github.com/tiagompalte/golang-clean-arch-template/internal/pkg/errors"
	"github.com/tiagompalte/golang-clean-arch-template/pkg/crypto"
	"github.com/tiagompalte/golang-clean-arch-template/pkg/errors"
)

type ValidateUserPasswordUseCase interface {
	Execute(ctx context.Context, input ValidateUserPasswordInput) (entity.User, error)
}

type ValidateUserPasswordInput struct {
	Email    string
	Password string
}

func (i ValidateUserPasswordInput) Validate() error {
	aggrErr := errors.NewAggregatedError()
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

type ValidateUserPasswordUseCaseImpl struct {
	userRepository repository.UserRepository
	crypto         crypto.Crypto
}

func NewValidateUserPasswordUseCaseImpl(userRepository repository.UserRepository, crypto crypto.Crypto) ValidateUserPasswordUseCase {
	return ValidateUserPasswordUseCaseImpl{
		userRepository: userRepository,
		crypto:         crypto,
	}
}

func (u ValidateUserPasswordUseCaseImpl) Execute(ctx context.Context, input ValidateUserPasswordInput) (entity.User, error) {
	err := input.Validate()
	if err != nil {
		return entity.User{}, errors.Wrap(err)
	}

	passEncrypted, err := u.userRepository.GetPassEncryptedByEmail(ctx, input.Email)
	if err != nil {
		return entity.User{}, errors.Wrap(errPkg.NewInvalidLoginError())
	}

	isValid, err := u.crypto.VerifyHash(ctx, input.Password, passEncrypted)
	if err != nil {
		return entity.User{}, errors.Wrap(err)
	}

	if !isValid {
		return entity.User{}, errors.Wrap(errPkg.NewInvalidLoginError())
	}

	user, err := u.userRepository.FindByEmail(ctx, input.Email)
	if err != nil {
		return entity.User{}, errors.Wrap(err)
	}

	return user, nil
}
