package usecase

import (
	"context"

	"github.com/tiagompalte/golang-clean-arch-template/internal/app/entity"
	errPkg "github.com/tiagompalte/golang-clean-arch-template/internal/pkg/errors"
	"github.com/tiagompalte/golang-clean-arch-template/internal/pkg/infra/uow"
	"github.com/tiagompalte/golang-clean-arch-template/pkg/errors"
)

type CreateTaskUseCase interface {
	Execute(ctx context.Context, input CreateTaskInput) (entity.Task, error)
}

type CreateTaskInput struct {
	Name         string
	Description  string
	CategoryName string
	UserID       uint32
}

func (i CreateTaskInput) Validate() error {
	aggrErr := errors.NewAggregatedError()
	if i.Name == "" {
		aggrErr.Add(errPkg.NewEmptyParameterError("name"))
	}
	if i.Description == "" {
		aggrErr.Add(errPkg.NewEmptyParameterError("description"))
	}
	if i.CategoryName == "" {
		aggrErr.Add(errPkg.NewEmptyParameterError("category"))
	}
	if i.UserID == 0 {
		aggrErr.Add(errPkg.NewEmptyParameterError("user_id"))
	}

	if aggrErr.Len() > 0 {
		return errors.Wrap(aggrErr)
	}

	return nil
}

type CreateTaskUseCaseImpl struct {
	uow uow.Uow
}

func NewCreateTaskUseCaseImpl(uow uow.Uow) CreateTaskUseCase {
	return CreateTaskUseCaseImpl{
		uow: uow,
	}
}

func (u CreateTaskUseCaseImpl) Execute(ctx context.Context, input CreateTaskInput) (entity.Task, error) {
	err := input.Validate()
	if err != nil {
		return entity.Task{}, errors.Wrap(err)
	}

	var task entity.Task
	err = u.uow.Do(ctx, func(uow *uow.Uow) error {
		categoryNew := entity.Category{
			Name:   input.CategoryName,
			UserID: input.UserID,
		}

		category, err := uow.Repository().Category().FindBySlugAndUserID(ctx, categoryNew.GetSlug(), categoryNew.UserID)
		if err != nil && !errors.IsAppError(err, errors.ErrorCodeNotFound) {
			return errors.Wrap(err)
		} else if errors.IsAppError(err, errors.ErrorCodeNotFound) {
			category.ID, err = uow.Repository().Category().Insert(ctx, categoryNew)
			if err != nil {
				return errors.Wrap(err)
			}
		}

		id, err := uow.Repository().Task().Insert(ctx, entity.Task{
			Name:        input.Name,
			Description: input.Description,
			Category:    category,
			UserID:      input.UserID,
		})
		if err != nil {
			return errors.Wrap(err)
		}

		task, err = uow.Repository().Task().FindByID(ctx, id)
		if err != nil {
			return errors.Wrap(err)
		}

		return nil
	})
	if err != nil {
		return entity.Task{}, errors.Wrap(err)
	}

	return task, nil
}
