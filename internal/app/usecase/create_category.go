package usecase

import (
	"context"

	"github.com/tiagompalte/golang-clean-arch-template/internal/app/entity"
	"github.com/tiagompalte/golang-clean-arch-template/internal/app/repository"
	errPkg "github.com/tiagompalte/golang-clean-arch-template/internal/pkg/errors"
	"github.com/tiagompalte/golang-clean-arch-template/pkg/errors"
)

type CreateCategoryUseCase interface {
	Execute(ctx context.Context, input CreateCategoryInput) (entity.Category, error)
}

type CreateCategoryInput struct {
	Name   string
	UserID uint32
}

func (i CreateCategoryInput) Validate() error {
	aggrErr := errors.NewAggregatedError()
	if i.Name == "" {
		aggrErr.Add(errPkg.NewEmptyParameterError("name"))
	}
	if i.UserID == 0 {
		aggrErr.Add(errPkg.NewEmptyParameterError("user_id"))
	}

	if aggrErr.Len() > 0 {
		return errors.Wrap(aggrErr)
	}

	return nil
}

type CreateCategoryUseCaseImpl struct {
	categoryRepository repository.CategoryRepository
}

func NewCreateCategoryUseCaseImpl(categoryRepository repository.CategoryRepository) CreateCategoryUseCase {
	return CreateCategoryUseCaseImpl{
		categoryRepository: categoryRepository,
	}
}

func (u CreateCategoryUseCaseImpl) Execute(ctx context.Context, input CreateCategoryInput) (entity.Category, error) {
	err := input.Validate()
	if err != nil {
		return entity.Category{}, errors.Wrap(err)
	}

	categoryNew := entity.Category{
		Name:   input.Name,
		UserID: input.UserID,
	}

	category, err := u.categoryRepository.FindBySlugAndUserID(ctx, categoryNew.GetSlug(), categoryNew.UserID)
	if err != nil && !errors.IsAppError(err, errors.ErrorCodeNotFound) {
		return entity.Category{}, errors.Wrap(err)
	} else if errors.IsAppError(err, errors.ErrorCodeNotFound) {
		categoryID, err := u.categoryRepository.Insert(ctx, categoryNew)
		if err != nil {
			return entity.Category{}, errors.Wrap(err)
		}

		category, err = u.categoryRepository.FindByID(ctx, categoryID)
		if err != nil {
			return entity.Category{}, errors.Wrap(err)
		}
	}

	return category, nil
}
