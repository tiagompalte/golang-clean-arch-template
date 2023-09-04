package usecase

import (
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	NewCreateTaskImpl,
	NewFindAllCategoryImpl,
	NewFindAllTaskImpl,
	NewFindOneTaskImpl,
	NewUpdateTaskDoneImpl,
	NewUpdateTaskUndoneImpl,
	NewDeleteTaskImpl,
	NewHealthCheckImpl,
)
