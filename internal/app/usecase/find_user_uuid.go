package usecase

import (
	"context"

	"github.com/tiagompalte/golang-clean-arch-template/internal/app/protocols"
	"github.com/tiagompalte/golang-clean-arch-template/pkg/errors"
)

type FindUserUUIDUseCase interface {
	Execute(ctx context.Context, uuid string) (FindUserUUIDOutput, error)
}

type FindUserUUIDOutput struct {
	ID      uint32
	Version uint32
	UUID    string
	Name    string
	Email   string
}

type FindUserUUIDUseCaseImpl struct {
	userRepository protocols.UserRepository
}

func NewFindUserUUIDUseCaseImpl(userRepository protocols.UserRepository) FindUserUUIDUseCase {
	return FindUserUUIDUseCaseImpl{
		userRepository: userRepository,
	}
}

func (u FindUserUUIDUseCaseImpl) Execute(ctx context.Context, uuid string) (FindUserUUIDOutput, error) {
	user, err := u.userRepository.FindByUUID(ctx, uuid)
	if err != nil {
		return FindUserUUIDOutput{}, errors.Wrap(err)
	}

	var output FindUserUUIDOutput
	output.ID = user.ID
	output.UUID = user.UUID
	output.Version = user.Version
	output.Name = user.Name
	output.Email = user.Email

	return output, nil
}
