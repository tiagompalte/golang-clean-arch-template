package usecase

import (
	"context"

	"github.com/tiagompalte/golang-clean-arch-template/internal/app/entity"
	"github.com/tiagompalte/golang-clean-arch-template/internal/pkg/infra/data"
	"github.com/tiagompalte/golang-clean-arch-template/internal/pkg/infra/uow"
	"github.com/tiagompalte/golang-clean-arch-template/pkg/errors"
	"github.com/tiagompalte/golang-clean-arch-template/pkg/repository"
)

type CreateTaskUseCase interface {
	Execute(ctx context.Context, input CreateTaskInput) (CreateTaskOutput, error)
}

type CreateTaskInput struct {
	Name         string
	Description  string
	CategoryName string
	UserID       uint32
}

type CreateTaskOutput struct {
	UUID         string
	Name         string
	Description  string
	CategorySlug string
	CategoryName string
	Done         bool
}

type CreateTaskUseCaseImpl struct {
	uow uow.Uow[repository.ConnectorSql, data.SqlManager]
}

func NewCreateTaskUseCaseImpl(uow uow.Uow[repository.ConnectorSql, data.SqlManager]) CreateTaskUseCase {
	return CreateTaskUseCaseImpl{
		uow: uow,
	}
}

func (u CreateTaskUseCaseImpl) Execute(ctx context.Context, input CreateTaskInput) (CreateTaskOutput, error) {
	categoryNew := entity.Category{
		Name:   input.CategoryName,
		UserID: input.UserID,
	}

	err := categoryNew.ValidateNew()
	if err != nil {
		return CreateTaskOutput{}, errors.Wrap(err)
	}

	task := entity.Task{
		Name:        input.Name,
		Description: input.Description,
		UserID:      input.UserID,
	}

	err = task.ValidateNew()
	if err != nil {
		return CreateTaskOutput{}, errors.Wrap(err)
	}

	err = u.uow.Do(ctx, func(data data.Manager[data.SqlManager]) error {
		category, err := data.Repository().Category().FindBySlugAndUserID(ctx, categoryNew.GetSlug(), categoryNew.UserID)
		if err != nil && !errors.IsAppError(err, errors.ErrorCodeNotFound) {
			return errors.Wrap(err)
		} else if errors.IsAppError(err, errors.ErrorCodeNotFound) {
			category.ID, err = data.Repository().Category().Insert(ctx, categoryNew)
			if err != nil {
				return errors.Wrap(err)
			}
		}

		task.Category = category

		id, err := data.Repository().Task().Insert(ctx, task)
		if err != nil {
			return errors.Wrap(err)
		}

		task, err = data.Repository().Task().FindByID(ctx, id)
		if err != nil {
			return errors.Wrap(err)
		}

		return nil
	})
	if err != nil {
		return CreateTaskOutput{}, errors.Wrap(err)
	}

	return CreateTaskOutput{
		UUID:         task.UUID,
		Name:         task.Name,
		Description:  task.Description,
		CategorySlug: task.Category.GetSlug(),
		CategoryName: task.Category.Name,
		Done:         task.Done,
	}, nil
}
