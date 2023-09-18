package usecase

import (
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	NewCreateTaskUseCaseImpl,
	NewFindAllCategoryUseCaseImpl,
	NewFindAllTaskUseCaseImpl,
	NewFindOneTaskUseCaseImpl,
	NewUpdateTaskDoneUseCaseImpl,
	NewUpdateTaskUndoneUseCaseImpl,
	NewDeleteTaskUseCaseImpl,
	NewHealthCheckUseCaseImpl,
	NewCreateUserUseCaseImpl,
	NewValidateUserPasswordUseCaseImpl,
	NewGenerateUserTokenUseCaseImpl,
	NewFindUserUUIDUseCaseImpl,
)
