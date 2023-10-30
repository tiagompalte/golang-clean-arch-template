package usecase

import (
	"context"

	"github.com/tiagompalte/golang-clean-arch-template/internal/app/repository"
	pkgErrors "github.com/tiagompalte/golang-clean-arch-template/internal/pkg/errors"
	"github.com/tiagompalte/golang-clean-arch-template/pkg/errors"
)

type DeleteTaskUseCase interface {
	Execute(ctx context.Context, input DeleteTaskUseCaseInput) error
}

type DeleteTaskUseCaseInput struct {
	UUID   string
	UserID uint32
}

type DeleteTaskUseCaseImpl struct {
	taskRepository repository.TaskRepository
}

func NewDeleteTaskUseCaseImpl(taskRepository repository.TaskRepository) DeleteTaskUseCase {
	return DeleteTaskUseCaseImpl{
		taskRepository: taskRepository,
	}
}

func (u DeleteTaskUseCaseImpl) Execute(ctx context.Context, input DeleteTaskUseCaseInput) error {
	task, err := u.taskRepository.FindByUUID(ctx, input.UUID)
	if err != nil {
		return errors.Wrap(err)
	}

	if task.UserID != input.UserID {
		return errors.Wrap(pkgErrors.NewInvalidUserError())
	}

	err = u.taskRepository.DeleteByID(ctx, task.ID)
	if err != nil {
		return errors.Wrap(err)
	}

	return nil
}
