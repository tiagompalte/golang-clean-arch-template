package routes

import (
	"github.com/tiagompalte/golang-clean-arch-template/application"
	"github.com/tiagompalte/golang-clean-arch-template/internal/pkg/server/constant"
	"github.com/tiagompalte/golang-clean-arch-template/internal/pkg/server/handler"
	"github.com/tiagompalte/golang-clean-arch-template/internal/pkg/server/middleware"
	"github.com/tiagompalte/golang-clean-arch-template/pkg/server"
)

func CreateGroupTaskV1(app application.App) server.GroupRoute {
	routes := []server.Route{
		{
			Path:    "/",
			Method:  server.RouteMethodPost,
			Handler: handler.CreateTaskHandler(app.UseCase().CreateTaskUseCase),
		},
		{
			Path:    "/",
			Method:  server.RouteMethodGet,
			Handler: handler.FindAllTaskHandler(app.UseCase().FindAllTaskUseCase),
		},
		{
			Path:    "/{uuid}",
			Method:  server.RouteMethodGet,
			Handler: handler.FindOneTaskHandler(app.UseCase().FindOneTaskUseCase),
		},
		{
			Path:    "/{uuid}/done",
			Method:  server.RouteMethodPut,
			Handler: handler.UpdateTaskDoneHandler(app.UseCase().UpdateTaskDoneUseCase),
		},
		{
			Path:    "/{uuid}/undone",
			Method:  server.RouteMethodPut,
			Handler: handler.UpdateTaskUndoneHandler(app.UseCase().UpdateTaskUndoneUseCase),
		},
		{
			Path:    "/{uuid}",
			Method:  server.RouteMethodDelete,
			Handler: handler.DeleteTaskHandler(app.UseCase().DeleteTaskUseCase),
		},
	}

	return server.GroupRoute{
		Path: "/tasks",
		Middlewares: []server.Middleware{
			middleware.ValidateExtractUserTokenMiddleware(constant.Authorization, app.Auth(), app.UseCase().FindUserUUIDUseCase),
		},
		Routes: routes,
	}
}
