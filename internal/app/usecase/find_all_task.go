package usecase

import (
	"context"

	"github.com/tiagompalte/golang-clean-arch-template/internal/app/protocols"
	"github.com/tiagompalte/golang-clean-arch-template/pkg/errors"
)

type FindAllTaskUseCase interface {
	Execute(ctx context.Context, userID uint32) ([]FindAllTaskOutput, error)
}

type FindAllTaskOutput struct {
	UUID         string
	Name         string
	Description  string
	CategorySlug string
	CategoryName string
	Done         bool
}

type FindAllTaskUseCaseImpl struct {
	taskRepository protocols.TaskRepository
}

func NewFindAllTaskUseCaseImpl(taskRepository protocols.TaskRepository) FindAllTaskUseCase {
	return FindAllTaskUseCaseImpl{
		taskRepository: taskRepository,
	}
}

func (u FindAllTaskUseCaseImpl) Execute(ctx context.Context, userID uint32) ([]FindAllTaskOutput, error) {
	list, err := u.taskRepository.FindByUserID(ctx, userID)
	if err != nil {
		return []FindAllTaskOutput{}, errors.Wrap(err)
	}

	output := make([]FindAllTaskOutput, len(list))
	for i := range list {
		output[i].UUID = list[i].UUID
		output[i].Name = list[i].Name
		output[i].Description = list[i].Description
		output[i].CategorySlug = list[i].Category.GetSlug()
		output[i].CategoryName = list[i].Category.Name
		output[i].Done = list[i].Done
	}

	return output, nil
}
