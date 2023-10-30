package usecase

import (
	"context"

	"github.com/tiagompalte/golang-clean-arch-template/internal/app/entity"
	"github.com/tiagompalte/golang-clean-arch-template/internal/app/repository"
	"github.com/tiagompalte/golang-clean-arch-template/pkg/errors"
)

type FindAllTaskUseCase interface {
	Execute(ctx context.Context, userID uint32) ([]entity.Task, error)
}

type FindAllTaskOutput struct {
	Items []entity.Task
}

type FindAllTaskUseCaseImpl struct {
	taskRepository repository.TaskRepository
}

func NewFindAllTaskUseCaseImpl(taskRepository repository.TaskRepository) FindAllTaskUseCase {
	return FindAllTaskUseCaseImpl{
		taskRepository: taskRepository,
	}
}

func (u FindAllTaskUseCaseImpl) Execute(ctx context.Context, userID uint32) ([]entity.Task, error) {
	list, err := u.taskRepository.FindByUserID(ctx, userID)
	if err != nil {
		return []entity.Task{}, errors.Wrap(err)
	}

	return list, nil
}
