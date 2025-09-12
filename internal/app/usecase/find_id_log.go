package usecase

import (
	"context"

	"github.com/tiagompalte/golang-clean-arch-template/internal/app/entity"
	"github.com/tiagompalte/golang-clean-arch-template/internal/app/protocols"
	"github.com/tiagompalte/golang-clean-arch-template/pkg/errors"
)

type FindByIDLogUseCase interface {
	Execute(ctx context.Context, id string) (entity.Log, error)
}

type FindByIDLogUseCaseImpl struct {
	logRepo protocols.LogRepository
}

func NewFindByIDLogUseCaseImpl(logRepo protocols.LogRepository) FindByIDLogUseCase {
	return FindByIDLogUseCaseImpl{logRepo: logRepo}
}

func (uc FindByIDLogUseCaseImpl) Execute(ctx context.Context, id string) (entity.Log, error) {
	log, err := uc.logRepo.FindByID(ctx, id)
	if err != nil {
		return entity.Log{}, errors.Wrap(err)
	}
	return log, nil
}
