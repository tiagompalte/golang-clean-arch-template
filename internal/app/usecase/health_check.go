package usecase

import (
	"context"

	"github.com/tiagompalte/golang-clean-arch-template/pkg/errors"
	"github.com/tiagompalte/golang-clean-arch-template/pkg/healthcheck"
)

type HealthCheckUseCase interface {
	Execute(ctx context.Context) error
}

type HealthCheckUseCaseImpl struct {
	implements []healthcheck.HealthCheck
}

func NewHealthCheckUseCaseImpl(implements []healthcheck.HealthCheck) HealthCheckUseCase {
	return HealthCheckUseCaseImpl{
		implements,
	}
}

func (u HealthCheckUseCaseImpl) Execute(ctx context.Context) error {
	for _, i := range u.implements {
		isHealthy, err := i.IsHealthy(ctx)
		if !isHealthy {
			return errors.Wrap(err)
		}
	}

	return nil
}
