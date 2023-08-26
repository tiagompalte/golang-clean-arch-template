package usecase

import (
	"context"

	"github.com/tiagompalte/golang-clean-arch-template/internal/app/repository"
	"github.com/tiagompalte/golang-clean-arch-template/pkg/errors"
	usecasePkg "github.com/tiagompalte/golang-clean-arch-template/pkg/usecase"
)

type FindOneTask usecasePkg.UseCase[FindOneTaskInput, FindOneTaskOutput]

type FindOneTaskInput struct {
	UUID string
}

type FindOneTaskOutput struct {
	UUID        string
	Name        string
	Description string
	Done        bool
}

type FindOneTaskImpl struct {
	taskRepository repository.TaskRepository
}

func NewFindOneTaskImpl(taskRepository repository.TaskRepository) FindOneTask {
	return FindOneTaskImpl{
		taskRepository: taskRepository,
	}
}

func (u FindOneTaskImpl) Execute(ctx context.Context, intput FindOneTaskInput) (FindOneTaskOutput, error) {
	task, err := u.taskRepository.FindByUUID(ctx, intput.UUID)
	if err != nil {
		return FindOneTaskOutput{}, errors.Wrap(err)
	}

	return FindOneTaskOutput{
		UUID:        task.UUID,
		Name:        task.Name,
		Description: task.Description,
		Done:        task.Done,
	}, nil
}
