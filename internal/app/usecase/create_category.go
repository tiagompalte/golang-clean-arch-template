package usecase

import (
	"context"

	"github.com/tiagompalte/golang-clean-arch-template/internal/app/entity"
	"github.com/tiagompalte/golang-clean-arch-template/internal/app/protocols"
	"github.com/tiagompalte/golang-clean-arch-template/pkg/errors"
)

type CreateCategoryUseCase interface {
	Execute(ctx context.Context, input CreateCategoryInput) (CreateCategoryOutput, error)
}

type CreateCategoryInput struct {
	Name   string
	UserID uint32
}

type CreateCategoryOutput struct {
	Slug string
	Name string
}

type CreateCategoryUseCaseImpl struct {
	categoryRepository protocols.CategoryRepository
}

func NewCreateCategoryUseCaseImpl(categoryRepository protocols.CategoryRepository) CreateCategoryUseCase {
	return CreateCategoryUseCaseImpl{
		categoryRepository: categoryRepository,
	}
}

func (u CreateCategoryUseCaseImpl) Execute(ctx context.Context, input CreateCategoryInput) (CreateCategoryOutput, error) {
	categoryNew := entity.Category{
		Name:   input.Name,
		UserID: input.UserID,
	}

	err := categoryNew.ValidateNew()
	if err != nil {
		return CreateCategoryOutput{}, errors.Wrap(err)
	}

	category, err := u.categoryRepository.FindBySlugAndUserID(ctx, categoryNew.GetSlug(), categoryNew.UserID)
	if err != nil && !errors.IsAppError(err, errors.ErrorCodeNotFound) {
		return CreateCategoryOutput{}, errors.Wrap(err)
	} else if errors.IsAppError(err, errors.ErrorCodeNotFound) {
		categoryID, err := u.categoryRepository.Insert(ctx, categoryNew)
		if err != nil {
			return CreateCategoryOutput{}, errors.Wrap(err)
		}

		category, err = u.categoryRepository.FindByID(ctx, categoryID)
		if err != nil {
			return CreateCategoryOutput{}, errors.Wrap(err)
		}
	}

	return CreateCategoryOutput{
		Slug: category.GetSlug(),
		Name: category.Name,
	}, nil
}
