package usecase

import (
	"context"

	"github.com/tiagompalte/golang-clean-arch-template/pkg/cache"
	"github.com/tiagompalte/golang-clean-arch-template/pkg/errors"
	pkgRepo "github.com/tiagompalte/golang-clean-arch-template/pkg/repository"
)

type HealthCheckUseCase interface {
	Execute(ctx context.Context) error
}

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

func (u HealthCheckUseCaseImpl) Execute(ctx context.Context) error {
	err := u.data.PingContext(ctx)
	if err != nil {
		return errors.Wrap(err)
	}

	err = u.cache.Ping(ctx)
	if err != nil {
		return errors.Wrap(err)
	}

	return nil
}
