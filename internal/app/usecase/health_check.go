package usecase

import (
	"context"

	"github.com/tiagompalte/golang-clean-arch-template/pkg/cache"
	"github.com/tiagompalte/golang-clean-arch-template/pkg/errors"
	pkgRepo "github.com/tiagompalte/golang-clean-arch-template/pkg/repository"
	usecasePkg "github.com/tiagompalte/golang-clean-arch-template/pkg/usecase"
)

type HealthCheck usecasePkg.UseCase[usecasePkg.Blank, usecasePkg.Blank]

type HealthCheckImpl struct {
	data  pkgRepo.DataManager
	cache cache.Cache
}

func NewHealthCheckImpl(data pkgRepo.DataManager, cache cache.Cache) HealthCheck {
	return HealthCheckImpl{
		data:  data,
		cache: cache,
	}
}

func (u HealthCheckImpl) Execute(ctx context.Context, _ usecasePkg.Blank) (usecasePkg.Blank, error) {
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
