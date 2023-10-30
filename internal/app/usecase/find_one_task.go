package usecase

import (
	"context"

	"github.com/tiagompalte/golang-clean-arch-template/internal/app/entity"
	"github.com/tiagompalte/golang-clean-arch-template/internal/app/repository"
	"github.com/tiagompalte/golang-clean-arch-template/pkg/errors"
)

type FindOneTaskUseCase interface {
	Execute(ctx context.Context, uuid string) (entity.Task, error)
}

type FindOneTaskUseCaseImpl struct {
	taskRepository repository.TaskRepository
}

func NewFindOneTaskUseCaseImpl(taskRepository repository.TaskRepository) FindOneTaskUseCase {
	return FindOneTaskUseCaseImpl{
		taskRepository: taskRepository,
	}
}

func (u FindOneTaskUseCaseImpl) Execute(ctx context.Context, uuid string) (entity.Task, error) {
	task, err := u.taskRepository.FindByUUID(ctx, uuid)
	if err != nil {
		return entity.Task{}, errors.Wrap(err)
	}

	return task, nil
}
