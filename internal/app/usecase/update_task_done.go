package usecase

import (
	"context"

	"github.com/tiagompalte/golang-clean-arch-template/internal/app/repository"
	pkgErrors "github.com/tiagompalte/golang-clean-arch-template/internal/pkg/errors"
	"github.com/tiagompalte/golang-clean-arch-template/pkg/errors"
	usecasePkg "github.com/tiagompalte/golang-clean-arch-template/pkg/usecase"
)

type UpdateTaskDoneUseCase usecasePkg.UseCase[UpdateTaskDoneUseCaseInput, usecasePkg.Blank]

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

func (u UpdateTaskDoneUseCaseImpl) Execute(ctx context.Context, input UpdateTaskDoneUseCaseInput) (usecasePkg.Blank, error) {
	task, err := u.taskRepository.FindByUUID(ctx, input.UUID)
	if err != nil {
		return usecasePkg.Blank{}, errors.Wrap(err)
	}

	if input.UserID != task.UserID {
		return usecasePkg.Blank{}, errors.Wrap(pkgErrors.NewInvalidUserError())
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
