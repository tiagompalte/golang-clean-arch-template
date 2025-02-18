package routes

import (
	"github.com/tiagompalte/golang-clean-arch-template/application"
	"github.com/tiagompalte/golang-clean-arch-template/internal/pkg/server/handler"
	"github.com/tiagompalte/golang-clean-arch-template/pkg/server"
)

func CreateRouteV1(app application.App) server.GroupRoute {
	return server.GroupRoute{
		Path: "/v1",
		GroupRoutes: []server.GroupRoute{
			CreateGroupTaskV1(app),
			CreateGroupCategoryV1(app),
			CreateGroupCurrentUserV1(app),
		},
		Routes: []server.Route{
			{
				Method:  server.RouteMethodPost,
				Path:    "/signup",
				Handler: handler.SignupHandler(app.UseCase().CreateUserUseCase, app.UseCase().GenerateUserTokenUseCase),
			},
			{
				Method:  server.RouteMethodPost,
				Path:    "/signin",
				Handler: handler.SigninHandler(app.UseCase().ValidateUserPasswordUseCase, app.UseCase().GenerateUserTokenUseCase),
			},
		},
	}
}
