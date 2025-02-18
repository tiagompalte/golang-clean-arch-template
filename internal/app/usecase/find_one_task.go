package usecase

import (
	"context"

	"github.com/tiagompalte/golang-clean-arch-template/internal/app/protocols"
	"github.com/tiagompalte/golang-clean-arch-template/pkg/errors"
)

type FindOneTaskUseCase interface {
	Execute(ctx context.Context, input FindOneTaskInput) (FindOneTaskOutput, error)
}

type FindOneTaskInput struct {
	TaskUUID string
	UserID   uint32
}

type FindOneTaskOutput struct {
	UUID         string
	Name         string
	Description  string
	CategorySlug string
	CategoryName string
	Done         bool
}

type FindOneTaskUseCaseImpl struct {
	taskRepository protocols.TaskRepository
}

func NewFindOneTaskUseCaseImpl(taskRepository protocols.TaskRepository) FindOneTaskUseCase {
	return FindOneTaskUseCaseImpl{
		taskRepository: taskRepository,
	}
}

func (u FindOneTaskUseCaseImpl) Execute(ctx context.Context, input FindOneTaskInput) (FindOneTaskOutput, error) {
	task, err := u.taskRepository.FindByUUID(ctx, input.TaskUUID)
	if err != nil {
		return FindOneTaskOutput{}, errors.Wrap(err)
	}

	if task.UserID != input.UserID {
		return FindOneTaskOutput{}, errors.Wrap(errors.NewInvalidUserError())
	}

	var output FindOneTaskOutput
	output.UUID = task.UUID
	output.Name = task.Name
	output.Description = task.Description
	output.CategorySlug = task.Category.GetSlug()
	output.CategoryName = task.Category.Name
	output.Done = task.Done

	return output, nil
}
