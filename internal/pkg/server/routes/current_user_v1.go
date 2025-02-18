package routes

import (
	"github.com/tiagompalte/golang-clean-arch-template/application"
	"github.com/tiagompalte/golang-clean-arch-template/internal/pkg/server/constant"
	"github.com/tiagompalte/golang-clean-arch-template/internal/pkg/server/handler"
	"github.com/tiagompalte/golang-clean-arch-template/internal/pkg/server/middleware"
	"github.com/tiagompalte/golang-clean-arch-template/pkg/server"
)

func CreateGroupCurrentUserV1(app application.App) server.GroupRoute {
	routes := []server.Route{
		{
			Path:    "/",
			Method:  server.RouteMethodGet,
			Handler: handler.FindUserLoggedHandler(),
		},
		{
			Path:    "/update-name",
			Method:  server.RouteMethodPut,
			Handler: handler.UpdateUserNameHandler(app.UseCase().UpdateUserNameUseCase),
		},
	}

	return server.GroupRoute{
		Path: "/current/user",
		Middlewares: []server.Middleware{
			middleware.ValidateExtractUserTokenMiddleware(constant.Authorization, app.Auth(), app.UseCase().FindUserUUIDUseCase),
		},
		Routes: routes,
	}
}
