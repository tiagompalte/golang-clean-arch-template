package usecase

import (
	"context"

	"github.com/tiagompalte/golang-clean-arch-template/pkg/errors"
	"github.com/tiagompalte/golang-clean-arch-template/pkg/healthcheck"
	"golang.org/x/sync/errgroup"
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
	group := errgroup.Group{}
	group.SetLimit(len(u.implements))

	for _, i := range u.implements {
		group.Go(func() error {
			_, err := i.IsHealthy(ctx)
			if err != nil {
				return errors.Wrap(err)
			}
			return nil
		})
	}

	err := group.Wait()
	if err != nil {
		return errors.Wrap(err)
	}

	return nil
}
