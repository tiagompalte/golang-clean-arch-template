package usecase

import (
	"context"

	"github.com/tiagompalte/golang-clean-arch-template/internal/app/repository"
	"github.com/tiagompalte/golang-clean-arch-template/pkg/errors"
	usecasePkg "github.com/tiagompalte/golang-clean-arch-template/pkg/usecase"
)

type DeleteTask usecasePkg.UseCase[string, usecasePkg.Blank]

type DeleteTaskImpl struct {
	taskRepository repository.TaskRepository
}

func NewDeleteTaskImpl(taskRepository repository.TaskRepository) DeleteTask {
	return DeleteTaskImpl{
		taskRepository: taskRepository,
	}
}

func (u DeleteTaskImpl) Execute(ctx context.Context, uuid string) (usecasePkg.Blank, error) {
	task, err := u.taskRepository.FindByUUID(ctx, uuid)
	if err != nil {
		return usecasePkg.Blank{}, errors.Wrap(err)
	}

	err = u.taskRepository.Delete(ctx, task)
	if err != nil {
		return usecasePkg.Blank{}, errors.Wrap(err)
	}

	return usecasePkg.Blank{}, nil
}
