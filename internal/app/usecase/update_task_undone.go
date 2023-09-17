package usecase

import (
	"context"

	"github.com/tiagompalte/golang-clean-arch-template/internal/app/repository"
	"github.com/tiagompalte/golang-clean-arch-template/pkg/errors"
	usecasePkg "github.com/tiagompalte/golang-clean-arch-template/pkg/usecase"
)

type UpdateTaskUndoneUseCase usecasePkg.UseCase[string, usecasePkg.Blank]

type UpdateTaskUndoneUseCaseImpl struct {
	taskRepository repository.TaskRepository
}

func NewUpdateTaskUndoneUseCaseImpl(taskRepository repository.TaskRepository) UpdateTaskUndoneUseCase {
	return UpdateTaskUndoneUseCaseImpl{
		taskRepository: taskRepository,
	}
}

func (u UpdateTaskUndoneUseCaseImpl) Execute(ctx context.Context, uuid string) (usecasePkg.Blank, error) {
	task, err := u.taskRepository.FindByUUID(ctx, uuid)
	if err != nil {
		return usecasePkg.Blank{}, errors.Wrap(err)
	}

	if !task.Done {
		return usecasePkg.Blank{}, nil
	}

	task.Done = false
	err = u.taskRepository.UpdateDone(ctx, task)
	if err != nil {
		return usecasePkg.Blank{}, errors.Wrap(err)
	}

	return usecasePkg.Blank{}, nil
}
