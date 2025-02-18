package usecase

import (
	"context"

	"github.com/tiagompalte/golang-clean-arch-template/internal/app/protocols"
	"github.com/tiagompalte/golang-clean-arch-template/pkg/crypto"
	"github.com/tiagompalte/golang-clean-arch-template/pkg/errors"
)

type ValidateUserPasswordUseCase interface {
	Execute(ctx context.Context, input ValidateUserPasswordInput) (ValidateUserPasswordOutput, error)
}

type ValidateUserPasswordInput struct {
	Email    string
	Password string
}

type ValidateUserPasswordOutput struct {
	UUID  string
	Name  string
	Email string
}

type ValidateUserPasswordUseCaseImpl struct {
	userRepository protocols.UserRepository
	crypto         crypto.Crypto
}

func NewValidateUserPasswordUseCaseImpl(userRepository protocols.UserRepository, crypto crypto.Crypto) ValidateUserPasswordUseCase {
	return ValidateUserPasswordUseCaseImpl{
		userRepository: userRepository,
		crypto:         crypto,
	}
}

func (u ValidateUserPasswordUseCaseImpl) validateInput(input ValidateUserPasswordInput) error {
	aggrErr := errors.NewAggregatedError()

	if input.Email == "" {
		aggrErr.Add(errors.NewEmptyParameterError("email"))
	}
	if input.Password == "" {
		aggrErr.Add(errors.NewEmptyParameterError("password"))
	}

	if aggrErr.Len() > 0 {
		return errors.Wrap(aggrErr)
	}

	return nil
}

func (u ValidateUserPasswordUseCaseImpl) Execute(ctx context.Context, input ValidateUserPasswordInput) (ValidateUserPasswordOutput, error) {
	err := u.validateInput(input)
	if err != nil {
		return ValidateUserPasswordOutput{}, errors.Wrap(err)
	}

	passEncrypted, err := u.userRepository.GetPassEncryptedByEmail(ctx, input.Email)
	if err != nil {
		return ValidateUserPasswordOutput{}, errors.Wrap(errors.NewInvalidLoginError())
	}

	isValid, err := u.crypto.VerifyHash(ctx, input.Password, passEncrypted)
	if err != nil {
		return ValidateUserPasswordOutput{}, errors.Wrap(err)
	}

	if !isValid {
		return ValidateUserPasswordOutput{}, errors.Wrap(errors.NewInvalidLoginError())
	}

	user, err := u.userRepository.FindByEmail(ctx, input.Email)
	if err != nil {
		return ValidateUserPasswordOutput{}, errors.Wrap(err)
	}

	var output ValidateUserPasswordOutput
	output.UUID = user.UUID
	output.Name = user.Name
	output.Email = user.Email

	return output, nil
}
