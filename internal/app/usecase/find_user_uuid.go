package usecase

import (
	"context"

	"github.com/tiagompalte/golang-clean-arch-template/internal/app/entity"
	"github.com/tiagompalte/golang-clean-arch-template/internal/app/repository"
	"github.com/tiagompalte/golang-clean-arch-template/pkg/errors"
)

type FindUserUUIDUseCase interface {
	Execute(ctx context.Context, uuid string) (entity.User, error)
}

type FindUserUUIDUseCaseImpl struct {
	userRepository repository.UserRepository
}

func NewFindUserUUIDUseCaseImpl(userRepository repository.UserRepository) FindUserUUIDUseCase {
	return FindUserUUIDUseCaseImpl{
		userRepository: userRepository,
	}
}

func (u FindUserUUIDUseCaseImpl) Execute(ctx context.Context, uuid string) (entity.User, error) {
	user, err := u.userRepository.FindByUUID(ctx, uuid)
	if err != nil {
		return entity.User{}, errors.Wrap(err)
	}

	return user, nil
}
