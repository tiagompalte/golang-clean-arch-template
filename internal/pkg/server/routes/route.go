package routes

import (
	"github.com/tiagompalte/golang-clean-arch-template/application"
	"github.com/tiagompalte/golang-clean-arch-template/internal/pkg/server/handler"
	"github.com/tiagompalte/golang-clean-arch-template/pkg/server"
)

func CreateRoute(app application.App) []server.GroupRoute {
	return []server.GroupRoute{
		{
			Path: "/api",
			GroupRoutes: []server.GroupRoute{
				CreateRouteV1(app),
			},
			Routes: []server.Route{
				{
					Method:  server.RouteMethodGet,
					Path:    "/health-check",
					Handler: handler.HealthCheckHandler(app.UseCase().HealthCheckUseCase),
				},
			},
		},
	}
}
