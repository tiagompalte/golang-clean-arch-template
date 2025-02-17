package usecase

import (
	"github.com/google/wire"
	"github.com/tiagompalte/golang-clean-arch-template/pkg/cache"
	"github.com/tiagompalte/golang-clean-arch-template/pkg/healthcheck"
	"github.com/tiagompalte/golang-clean-arch-template/pkg/repository"
)

var ProviderSet = wire.NewSet(
	NewCreateCategoryUseCaseImpl,
	NewCreateTaskUseCaseImpl,
	NewFindAllCategoryUseCaseImpl,
	NewFindAllTaskUseCaseImpl,
	NewFindOneTaskUseCaseImpl,
	NewUpdateTaskDoneUseCaseImpl,
	NewUpdateTaskUndoneUseCaseImpl,
	NewDeleteTaskUseCaseImpl,
	NewCreateUserUseCaseImpl,
	NewValidateUserPasswordUseCaseImpl,
	NewGenerateUserTokenUseCaseImpl,
	NewFindUserUUIDUseCaseImpl,
	ProviderHealthCheckUseCase,
	NewUpdateUserNameUseCaseImpl,
)

func ProviderHealthCheckUseCase(cache cache.Cache, dataSqlManager repository.DataSqlManager) HealthCheckUseCase {
	return NewHealthCheckUseCaseImpl([]healthcheck.HealthCheck{
		cache, dataSqlManager,
	})
}
