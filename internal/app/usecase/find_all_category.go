package usecase

import (
	"context"

	"github.com/tiagompalte/golang-clean-arch-template/internal/app/entity"
	"github.com/tiagompalte/golang-clean-arch-template/internal/app/repository"
	"github.com/tiagompalte/golang-clean-arch-template/pkg/errors"
	usecasePkg "github.com/tiagompalte/golang-clean-arch-template/pkg/usecase"
)

type FindAllCategoryUseCase usecasePkg.UseCase[uint32, FindAllCategoryOutput]

type FindAllCategoryOutput struct {
	Items []entity.Category
}

type FindAllCategoryUseCaseImpl struct {
	categoryRepository repository.CategoryRepository
}

func NewFindAllCategoryUseCaseImpl(categoryRepository repository.CategoryRepository) FindAllCategoryUseCase {
	return FindAllCategoryUseCaseImpl{
		categoryRepository: categoryRepository,
	}
}

func (u FindAllCategoryUseCaseImpl) Execute(ctx context.Context, userID uint32) (FindAllCategoryOutput, error) {
	list, err := u.categoryRepository.FindByUserID(ctx, userID)
	if err != nil {
		return FindAllCategoryOutput{}, errors.Wrap(err)
	}

	return FindAllCategoryOutput{Items: list}, nil
}
