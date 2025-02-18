package routes

import (
	"github.com/tiagompalte/golang-clean-arch-template/application"
	"github.com/tiagompalte/golang-clean-arch-template/internal/pkg/server/constant"
	"github.com/tiagompalte/golang-clean-arch-template/internal/pkg/server/handler"
	"github.com/tiagompalte/golang-clean-arch-template/internal/pkg/server/middleware"
	"github.com/tiagompalte/golang-clean-arch-template/pkg/server"
)

func CreateGroupCategoryV1(app application.App) server.GroupRoute {
	routes := []server.Route{
		{
			Path:    "/",
			Method:  server.RouteMethodGet,
			Handler: handler.FindAllCategoryHandler(app.UseCase().FindAllCategoryUseCase),
		},
	}

	return server.GroupRoute{
		Path: "/categories",
		Middlewares: []server.Middleware{
			middleware.ValidateExtractUserTokenMiddleware(constant.Authorization, app.Auth(), app.UseCase().FindUserUUIDUseCase),
		},
		Routes: routes,
	}
}
