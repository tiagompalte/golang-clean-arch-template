// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package application

import (
	"github.com/tiagompalte/golang-clean-arch-template/internal/app/usecase"
	"github.com/tiagompalte/golang-clean-arch-template/internal/pkg/infra/data"
	"github.com/tiagompalte/golang-clean-arch-template/internal/pkg/infra/uow"
	"github.com/tiagompalte/golang-clean-arch-template/pkg/auth"
	"github.com/tiagompalte/golang-clean-arch-template/pkg/cache"
	"github.com/tiagompalte/golang-clean-arch-template/pkg/config"
	"github.com/tiagompalte/golang-clean-arch-template/pkg/crypto"
	"github.com/tiagompalte/golang-clean-arch-template/pkg/log"
	"github.com/tiagompalte/golang-clean-arch-template/pkg/repository"
	"github.com/tiagompalte/golang-clean-arch-template/pkg/server"
)

// Injectors from wire.go:

func Build() (App, error) {
	configsConfig := config.ProviderSet()
	serverServer := server.ProviderSet(configsConfig)
	connectorSql := repository.ProviderConnectorSqlSet(configsConfig)
	categoryRepository := data.NewCategoryRepository(connectorSql)
	createCategoryUseCase := usecase.NewCreateCategoryUseCaseImpl(categoryRepository)
	dataManager := repository.ProviderDataSqlManagerSet(configsConfig)
	uowUow := uow.NewUow(dataManager)
	createTaskUseCase := usecase.NewCreateTaskUseCaseImpl(uowUow)
	findAllCategoryUseCase := usecase.NewFindAllCategoryUseCaseImpl(categoryRepository)
	taskRepository := data.NewTaskRepository(connectorSql)
	findAllTaskUseCase := usecase.NewFindAllTaskUseCaseImpl(taskRepository)
	findOneTaskUseCase := usecase.NewFindOneTaskUseCaseImpl(taskRepository)
	updateTaskDoneUseCase := usecase.NewUpdateTaskDoneUseCaseImpl(taskRepository)
	updateTaskUndoneUseCase := usecase.NewUpdateTaskUndoneUseCaseImpl(taskRepository)
	deleteTaskUseCase := usecase.NewDeleteTaskUseCaseImpl(taskRepository)
	cacheCache := cache.ProviderSet(configsConfig)
	healthCheckUseCase := usecase.ProviderHealthCheckUseCase(cacheCache, dataManager)
	userRepository := data.NewUserRepository(connectorSql)
	cryptoCrypto := crypto.ProviderSet(configsConfig)
	createUserUseCase := usecase.NewCreateUserUseCaseImpl(userRepository, cryptoCrypto)
	validateUserPasswordUseCase := usecase.NewValidateUserPasswordUseCaseImpl(userRepository, cryptoCrypto)
	authAuth := auth.ProviderSet(configsConfig)
	generateUserTokenUseCase := usecase.NewGenerateUserTokenUseCaseImpl(authAuth)
	findUserUUIDUseCase := usecase.NewFindUserUUIDUseCaseImpl(userRepository)
	updateUserNameUseCase := usecase.NewUpdateUserNameUseCaseImpl(userRepository)
	useCase := usecase.UseCase{
		CreateCategoryUseCase:       createCategoryUseCase,
		CreateTaskUseCase:           createTaskUseCase,
		FindAllCategoryUseCase:      findAllCategoryUseCase,
		FindAllTaskUseCase:          findAllTaskUseCase,
		FindOneTaskUseCase:          findOneTaskUseCase,
		UpdateTaskDoneUseCase:       updateTaskDoneUseCase,
		UpdateTaskUndoneUseCase:     updateTaskUndoneUseCase,
		DeleteTaskUseCase:           deleteTaskUseCase,
		HealthCheckUseCase:          healthCheckUseCase,
		CreateUserUseCase:           createUserUseCase,
		ValidateUserPasswordUseCase: validateUserPasswordUseCase,
		GenerateUserTokenUseCase:    generateUserTokenUseCase,
		FindUserUUIDUseCase:         findUserUUIDUseCase,
		UpdateUserNameUseCase:       updateUserNameUseCase,
	}
	logLog := log.ProviderSet()
	app := ProvideApplication(configsConfig, serverServer, useCase, authAuth, logLog)
	return app, nil
}
