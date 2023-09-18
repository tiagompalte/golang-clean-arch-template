package usecase

type UseCase struct {
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
}
