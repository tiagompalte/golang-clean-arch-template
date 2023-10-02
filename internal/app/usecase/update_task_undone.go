package usecase

import (
	"context"

	"github.com/tiagompalte/golang-clean-arch-template/internal/app/repository"
	pkgErrors "github.com/tiagompalte/golang-clean-arch-template/internal/pkg/errors"
	"github.com/tiagompalte/golang-clean-arch-template/pkg/errors"
	usecasePkg "github.com/tiagompalte/golang-clean-arch-template/pkg/usecase"
)

type UpdateTaskUndoneUseCase usecasePkg.UseCase[UpdateTaskUndoneUseCaseInput, usecasePkg.Blank]

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

func (u UpdateTaskUndoneUseCaseImpl) Execute(ctx context.Context, input UpdateTaskUndoneUseCaseInput) (usecasePkg.Blank, error) {
	task, err := u.taskRepository.FindByUUID(ctx, input.UUID)
	if err != nil {
		return usecasePkg.Blank{}, errors.Wrap(err)
	}

	if input.UserID != task.UserID {
		return usecasePkg.Blank{}, errors.Wrap(pkgErrors.NewInvalidUserError())
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
