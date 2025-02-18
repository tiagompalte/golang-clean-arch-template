package usecase

import (
	"context"

	"github.com/tiagompalte/golang-clean-arch-template/internal/app/protocols"
	"github.com/tiagompalte/golang-clean-arch-template/pkg/errors"
)

type FindAllCategoryUseCase interface {
	Execute(ctx context.Context, userID uint32) ([]FindAllCategoryOutput, error)
}

type FindAllCategoryOutput struct {
	Slug string
	Name string
}

type FindAllCategoryUseCaseImpl struct {
	categoryRepository protocols.CategoryRepository
}

func NewFindAllCategoryUseCaseImpl(categoryRepository protocols.CategoryRepository) FindAllCategoryUseCase {
	return FindAllCategoryUseCaseImpl{
		categoryRepository: categoryRepository,
	}
}

func (u FindAllCategoryUseCaseImpl) Execute(ctx context.Context, userID uint32) ([]FindAllCategoryOutput, error) {
	list, err := u.categoryRepository.FindByUserID(ctx, userID)
	if err != nil {
		return []FindAllCategoryOutput{}, errors.Wrap(err)
	}

	output := make([]FindAllCategoryOutput, len(list))
	for i := range list {
		output[i].Slug = list[i].GetSlug()
		output[i].Name = list[i].Name
	}

	return output, nil
}
