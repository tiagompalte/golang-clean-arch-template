package usecase

import (
	"context"

	"github.com/tiagompalte/golang-clean-arch-template/internal/app/repository"
	pkgErrors "github.com/tiagompalte/golang-clean-arch-template/internal/pkg/errors"
	"github.com/tiagompalte/golang-clean-arch-template/pkg/errors"
)

type UpdateTaskUndoneUseCase interface {
	Execute(ctx context.Context, input UpdateTaskUndoneUseCaseInput) error
}

type UpdateTaskUndoneUseCaseInput struct {
	UUID   string
	UserID uint32
}

type UpdateTaskUndoneUseCaseImpl struct {
	taskRepository repository.TaskRepository
}

func NewUpdateTaskUndoneUseCaseImpl(taskRepository repository.TaskRepository) UpdateTaskUndoneUseCase {
	return UpdateTaskUndoneUseCaseImpl{
		taskRepository: taskRepository,
	}
}

func (u UpdateTaskUndoneUseCaseImpl) Execute(ctx context.Context, input UpdateTaskUndoneUseCaseInput) error {
	task, err := u.taskRepository.FindByUUID(ctx, input.UUID)
	if err != nil {
		return errors.Wrap(err)
	}

	if input.UserID != task.UserID {
		return errors.Wrap(pkgErrors.NewInvalidUserError())
	}

	if !task.Done {
		return nil
	}

	task.Done = false
	err = u.taskRepository.UpdateDone(ctx, task)
	if err != nil {
		return errors.Wrap(err)
	}

	return nil
}
