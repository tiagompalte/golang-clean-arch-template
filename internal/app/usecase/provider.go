package usecase

import (
	"github.com/google/wire"
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
	NewHealthCheckUseCaseImpl,
	NewCreateUserUseCaseImpl,
	NewValidateUserPasswordUseCaseImpl,
	NewGenerateUserTokenUseCaseImpl,
	NewFindUserUUIDUseCaseImpl,
)
