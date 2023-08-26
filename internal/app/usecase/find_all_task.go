package usecase

import (
	"context"

	"github.com/tiagompalte/golang-clean-arch-template/internal/app/repository"
	"github.com/tiagompalte/golang-clean-arch-template/pkg/errors"
	usecasePkg "github.com/tiagompalte/golang-clean-arch-template/pkg/usecase"
)

type FindAllTask usecasePkg.UseCase[usecasePkg.Blank, FindAllTaskOutput]

type ItemFindAllTaskOutput struct {
	UUID        string
	Name        string
	Description string
	Done        bool
}

type FindAllTaskOutput struct {
	Items []ItemFindAllTaskOutput
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

	items := make([]ItemFindAllTaskOutput, 0, len(list))
	for _, task := range list {
		items = append(items, ItemFindAllTaskOutput{
			UUID:        task.UUID,
			Name:        task.Name,
			Description: task.Description,
			Done:        task.Done,
		})
	}

	return FindAllTaskOutput{Items: items}, nil
}
