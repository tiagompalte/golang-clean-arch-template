package usecase

import (
	"context"

	"github.com/tiagompalte/golang-clean-arch-template/internal/app/repository"
	pkgErrors "github.com/tiagompalte/golang-clean-arch-template/internal/pkg/errors"
	"github.com/tiagompalte/golang-clean-arch-template/pkg/errors"
)

type UpdateTaskDoneUseCase interface {
	Execute(ctx context.Context, input UpdateTaskDoneUseCaseInput) error
}

type UpdateTaskDoneUseCaseInput struct {
	UUID   string
	UserID uint32
}

type UpdateTaskDoneUseCaseImpl struct {
	taskRepository repository.TaskRepository
}

func NewUpdateTaskDoneUseCaseImpl(taskRepository repository.TaskRepository) UpdateTaskDoneUseCase {
	return UpdateTaskDoneUseCaseImpl{
		taskRepository: taskRepository,
	}
}

func (u UpdateTaskDoneUseCaseImpl) Execute(ctx context.Context, input UpdateTaskDoneUseCaseInput) error {
	task, err := u.taskRepository.FindByUUID(ctx, input.UUID)
	if err != nil {
		return errors.Wrap(err)
	}

	if input.UserID != task.UserID {
		return errors.Wrap(pkgErrors.NewInvalidUserError())
	}

	if task.Done {
		return nil
	}

	task.Done = true
	err = u.taskRepository.UpdateDone(ctx, task)
	if err != nil {
		return errors.Wrap(err)
	}

	return nil
}
