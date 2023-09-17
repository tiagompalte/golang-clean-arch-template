package usecase

import (
	"context"

	"github.com/tiagompalte/golang-clean-arch-template/internal/app/repository"
	"github.com/tiagompalte/golang-clean-arch-template/pkg/errors"
	usecasePkg "github.com/tiagompalte/golang-clean-arch-template/pkg/usecase"
)

type UpdateTaskDoneUseCase usecasePkg.UseCase[string, usecasePkg.Blank]

type UpdateTaskDoneUseCaseImpl struct {
	taskRepository repository.TaskRepository
}

func NewUpdateTaskDoneUseCaseImpl(taskRepository repository.TaskRepository) UpdateTaskDoneUseCase {
	return UpdateTaskDoneUseCaseImpl{
		taskRepository: taskRepository,
	}
}

func (u UpdateTaskDoneUseCaseImpl) Execute(ctx context.Context, uuid string) (usecasePkg.Blank, error) {
	task, err := u.taskRepository.FindByUUID(ctx, uuid)
	if err != nil {
		return usecasePkg.Blank{}, errors.Wrap(err)
	}

	if task.Done {
		return usecasePkg.Blank{}, nil
	}

	task.Done = true
	err = u.taskRepository.UpdateDone(ctx, task)
	if err != nil {
		return usecasePkg.Blank{}, errors.Wrap(err)
	}

	return usecasePkg.Blank{}, nil
}
