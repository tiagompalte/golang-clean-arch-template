package usecase

import (
	"context"

	"github.com/tiagompalte/golang-clean-arch-template/internal/app/entity"
	"github.com/tiagompalte/golang-clean-arch-template/internal/app/repository"
	"github.com/tiagompalte/golang-clean-arch-template/pkg/errors"
	usecasePkg "github.com/tiagompalte/golang-clean-arch-template/pkg/usecase"
)

type FindOneTask usecasePkg.UseCase[string, entity.Task]

type FindOneTaskImpl struct {
	taskRepository repository.TaskRepository
}

func NewFindOneTaskImpl(taskRepository repository.TaskRepository) FindOneTask {
	return FindOneTaskImpl{
		taskRepository: taskRepository,
	}
}

func (u FindOneTaskImpl) Execute(ctx context.Context, uuid string) (entity.Task, error) {
	task, err := u.taskRepository.FindByUUID(ctx, uuid)
	if err != nil {
		return entity.Task{}, errors.Wrap(err)
	}

	return task, nil
}
