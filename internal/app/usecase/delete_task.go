package usecase

import (
	"context"

	"github.com/tiagompalte/golang-clean-arch-template/internal/app/repository"
	pkgErrors "github.com/tiagompalte/golang-clean-arch-template/internal/pkg/errors"
	"github.com/tiagompalte/golang-clean-arch-template/pkg/errors"
	usecasePkg "github.com/tiagompalte/golang-clean-arch-template/pkg/usecase"
)

type DeleteTaskUseCase usecasePkg.UseCase[DeleteTaskUseCaseInput, usecasePkg.Blank]

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

func (u DeleteTaskUseCaseImpl) Execute(ctx context.Context, input DeleteTaskUseCaseInput) (usecasePkg.Blank, error) {
	task, err := u.taskRepository.FindByUUID(ctx, input.UUID)
	if err != nil {
		return usecasePkg.Blank{}, errors.Wrap(err)
	}

	if task.UserID != input.UserID {
		return usecasePkg.Blank{}, errors.Wrap(pkgErrors.NewInvalidUserError())
	}

	err = u.taskRepository.Delete(ctx, task)
	if err != nil {
		return usecasePkg.Blank{}, errors.Wrap(err)
	}

	return usecasePkg.Blank{}, nil
}
