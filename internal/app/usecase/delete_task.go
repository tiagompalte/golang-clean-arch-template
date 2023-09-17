package usecase

import (
	"context"

	"github.com/tiagompalte/golang-clean-arch-template/internal/app/repository"
	"github.com/tiagompalte/golang-clean-arch-template/pkg/errors"
	usecasePkg "github.com/tiagompalte/golang-clean-arch-template/pkg/usecase"
)

type DeleteTaskUseCase usecasePkg.UseCase[string, usecasePkg.Blank]

type DeleteTaskUseCaseImpl struct {
	taskRepository repository.TaskRepository
}

func NewDeleteTaskUseCaseImpl(taskRepository repository.TaskRepository) DeleteTaskUseCase {
	return DeleteTaskUseCaseImpl{
		taskRepository: taskRepository,
	}
}

func (u DeleteTaskUseCaseImpl) Execute(ctx context.Context, uuid string) (usecasePkg.Blank, error) {
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
