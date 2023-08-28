package usecase

import (
	"context"

	"github.com/tiagompalte/golang-clean-arch-template/internal/app/repository"
	"github.com/tiagompalte/golang-clean-arch-template/pkg/errors"
	usecasePkg "github.com/tiagompalte/golang-clean-arch-template/pkg/usecase"
)

type UpdateTaskUndone usecasePkg.UseCase[string, usecasePkg.Blank]

type UpdateTaskUndoneImpl struct {
	taskRepository repository.TaskRepository
}

func NewUpdateTaskUndoneImpl(taskRepository repository.TaskRepository) UpdateTaskUndone {
	return UpdateTaskUndoneImpl{
		taskRepository: taskRepository,
	}
}

func (u UpdateTaskUndoneImpl) Execute(ctx context.Context, uuid string) (usecasePkg.Blank, error) {
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
