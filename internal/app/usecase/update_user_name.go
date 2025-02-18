package usecase

import (
	"context"

	"github.com/tiagompalte/golang-clean-arch-template/internal/app/entity"
	"github.com/tiagompalte/golang-clean-arch-template/internal/app/protocols"
	"github.com/tiagompalte/golang-clean-arch-template/pkg/errors"
)

type UpdateUserNameUseCase interface {
	Execute(ctx context.Context, input UpdateUserNameInput) (UpdateUserNameOutput, error)
}

type UpdateUserNameInput struct {
	UserUUID    string
	NewUserName string
	Version     uint32
}

type UpdateUserNameOutput struct {
	UUID    string
	Name    string
	Email   string
	Version uint32
}

type UpdateUserNameUseCaseImpl struct {
	userRepository protocols.UserRepository
}

func NewUpdateUserNameUseCaseImpl(userRepository protocols.UserRepository) UpdateUserNameUseCase {
	return UpdateUserNameUseCaseImpl{
		userRepository: userRepository,
	}
}

func (u UpdateUserNameUseCaseImpl) validate(user entity.User) error {
	aggr := errors.NewAggregatedError()
	if user.UUID == "" {
		aggr.Add(errors.NewEmptyParameterError("user_uuid"))
	}
	if user.Name == "" {
		aggr.Add(errors.NewEmptyParameterError("user_name"))
	}
	return aggr.Return()
}

func (u UpdateUserNameUseCaseImpl) Execute(ctx context.Context, input UpdateUserNameInput) (UpdateUserNameOutput, error) {
	var user entity.User
	user.UUID = input.UserUUID
	user.Name = input.NewUserName
	user.Version = input.Version

	err := u.validate(user)
	if err != nil {
		return UpdateUserNameOutput{}, errors.Wrap(err)
	}

	userDatabase, err := u.userRepository.FindByUUID(ctx, input.UserUUID)
	if err != nil {
		return UpdateUserNameOutput{}, errors.Wrap(err)
	}
	user.ID = userDatabase.ID

	err = u.userRepository.UpdateName(ctx, user)
	if err != nil {
		return UpdateUserNameOutput{}, errors.Wrap(err)
	}

	user, err = u.userRepository.FindByID(ctx, user.ID)
	if err != nil {
		return UpdateUserNameOutput{}, errors.Wrap(err)
	}

	return UpdateUserNameOutput{
		UUID:    user.UUID,
		Name:    user.Name,
		Email:   user.Email,
		Version: user.Version,
	}, nil
}
