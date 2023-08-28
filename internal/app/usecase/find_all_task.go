package usecase

import (
	"context"

	"github.com/tiagompalte/golang-clean-arch-template/internal/app/entity"
	"github.com/tiagompalte/golang-clean-arch-template/internal/app/repository"
	"github.com/tiagompalte/golang-clean-arch-template/pkg/errors"
	usecasePkg "github.com/tiagompalte/golang-clean-arch-template/pkg/usecase"
)

type FindAllTask usecasePkg.UseCase[usecasePkg.Blank, FindAllTaskOutput]

type FindAllTaskOutput struct {
	Items []entity.Task
}

type FindAllTaskImpl struct {
	taskRepository repository.TaskRepository
}

func NewFindAllTaskImpl(taskRepository repository.TaskRepository) FindAllTask {
	return FindAllTaskImpl{
		taskRepository: taskRepository,
	}
}

func (u FindAllTaskImpl) Execute(ctx context.Context, _ usecasePkg.Blank) (FindAllTaskOutput, error) {
	list, err := u.taskRepository.FindAll(ctx)
	if err != nil {
		return FindAllTaskOutput{}, errors.Wrap(err)
	}

	return FindAllTaskOutput{Items: list}, nil
}
