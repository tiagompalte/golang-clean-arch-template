package usecase

type UseCase struct {
	CreateCategoryUseCase
	CreateLogUseCase
	CreateTaskUseCase
	FindAllCategoryUseCase
	FindAllTaskUseCase
	FindOneTaskUseCase
	UpdateTaskDoneUseCase
	UpdateTaskUndoneUseCase
	DeleteTaskUseCase
	HealthCheckUseCase
	CreateUserUseCase
	ValidateUserPasswordUseCase
	GenerateUserTokenUseCase
	FindUserUUIDUseCase
	UpdateUserNameUseCase
}
