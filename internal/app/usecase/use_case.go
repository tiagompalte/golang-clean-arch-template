package usecase

type UseCase struct {
	CreateTask       CreateTask
	FindAllCategory  FindAllCategory
	FindAllTask      FindAllTask
	FindOneTask      FindOneTask
	UpdateTaskDone   UpdateTaskDone
	UpdateTaskUndone UpdateTaskUndone
	DeleteTask       DeleteTask
	HealthCheck      HealthCheck
}
