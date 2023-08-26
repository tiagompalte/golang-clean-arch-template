package application

import (
	"github.com/tiagompalte/golang-clean-arch-template/configs"
	"github.com/tiagompalte/golang-clean-arch-template/internal/app/usecase"
	"github.com/tiagompalte/golang-clean-arch-template/pkg/logger"
	"github.com/tiagompalte/golang-clean-arch-template/pkg/server"
)

type App struct {
	config  configs.Config
	server  server.Server
	logger  logger.Logger
	useCase usecase.UseCase
}

func ProvideApplication(
	config configs.Config,
	server server.Server,
	logger logger.Logger,
	useCase usecase.UseCase,
) App {
	return App{
		config,
		server,
		logger,
		useCase,
	}
}

func (app App) Config() configs.Config {
	return app.config
}

func (app App) Server() server.Server {
	return app.server
}

func (app App) Logger() logger.Logger {
	return app.logger
}

func (app App) UseCase() usecase.UseCase {
	return app.useCase
}
