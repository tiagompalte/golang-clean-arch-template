package v1

import (
	"github.com/tiagompalte/golang-clean-arch-template/application"
	"github.com/tiagompalte/golang-clean-arch-template/internal/pkg/server/handler"
	"github.com/tiagompalte/golang-clean-arch-template/pkg/server"
)

func CreateRoute(app application.App) server.GroupRoute {
	return server.GroupRoute{
		Path: "/v1",
		GroupRoutes: []server.GroupRoute{
			CreateGroupTask(app),
			CreateGroupCategory(app),
		},
		Routes: []server.Route{
			{
				Method:  "POST",
				Path:    "/signup",
				Handler: handler.SignupHandler(app.UseCase().CreateUserUseCase, app.UseCase().GenerateUserTokenUseCase),
			},
			{
				Method:  "POST",
				Path:    "/signin",
				Handler: handler.SigninHandler(app.UseCase().ValidateUserPasswordUseCase, app.UseCase().GenerateUserTokenUseCase),
			},
		},
	}
}
