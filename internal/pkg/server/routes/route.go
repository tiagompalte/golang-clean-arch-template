package routes

import (
	"github.com/tiagompalte/golang-clean-arch-template/application"
	"github.com/tiagompalte/golang-clean-arch-template/internal/pkg/server/handler"
	v1 "github.com/tiagompalte/golang-clean-arch-template/internal/pkg/server/routes/v1"
	"github.com/tiagompalte/golang-clean-arch-template/pkg/server"
)

func CreateRoute(app application.App) []server.GroupRoute {
	return []server.GroupRoute{
		{
			Path: "/api",
			GroupRoutes: []server.GroupRoute{
				v1.CreateRoute(app),
			},
			Routes: []server.Route{
				{
					Method:  "GET",
					Path:    "/health-check",
					Handler: handler.HealthCheckHandler(app.UseCase().HealthCheckUseCase),
				},
			},
		},
	}
}
