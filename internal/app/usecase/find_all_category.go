package usecase

import (
	"context"

	"github.com/tiagompalte/golang-clean-arch-template/internal/app/entity"
	"github.com/tiagompalte/golang-clean-arch-template/internal/app/repository"
	"github.com/tiagompalte/golang-clean-arch-template/pkg/errors"
)

type FindAllCategoryUseCase interface {
	Execute(ctx context.Context, userID uint32) ([]entity.Category, error)
}

type FindAllCategoryUseCaseImpl struct {
	categoryRepository repository.CategoryRepository
}

func NewFindAllCategoryUseCaseImpl(categoryRepository repository.CategoryRepository) FindAllCategoryUseCase {
	return FindAllCategoryUseCaseImpl{
		categoryRepository: categoryRepository,
	}
}

func (u FindAllCategoryUseCaseImpl) Execute(ctx context.Context, userID uint32) ([]entity.Category, error) {
	list, err := u.categoryRepository.FindByUserID(ctx, userID)
	if err != nil {
		return []entity.Category{}, errors.Wrap(err)
	}

	return list, nil
}
