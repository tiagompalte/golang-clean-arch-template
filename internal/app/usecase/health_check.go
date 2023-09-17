package usecase

import (
	"context"

	"github.com/tiagompalte/golang-clean-arch-template/pkg/cache"
	"github.com/tiagompalte/golang-clean-arch-template/pkg/errors"
	pkgRepo "github.com/tiagompalte/golang-clean-arch-template/pkg/repository"
	usecasePkg "github.com/tiagompalte/golang-clean-arch-template/pkg/usecase"
)

type HealthCheckUseCase usecasePkg.UseCase[usecasePkg.Blank, usecasePkg.Blank]

type HealthCheckUseCaseImpl struct {
	data  pkgRepo.DataManager
	cache cache.Cache
}

func NewHealthCheckUseCaseImpl(data pkgRepo.DataManager, cache cache.Cache) HealthCheckUseCase {
	return HealthCheckUseCaseImpl{
		data:  data,
		cache: cache,
	}
}

func (u HealthCheckUseCaseImpl) Execute(ctx context.Context, _ usecasePkg.Blank) (usecasePkg.Blank, error) {
	err := u.data.PingContext(ctx)
	if err != nil {
		return usecasePkg.Blank{}, errors.Wrap(err)
	}

	err = u.cache.Ping(ctx)
	if err != nil {
		return usecasePkg.Blank{}, errors.Wrap(err)
	}

	return usecasePkg.Blank{}, nil
}
