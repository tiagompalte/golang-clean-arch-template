package routes

import (
	"github.com/tiagompalte/golang-clean-arch-template/application"
	v1 "github.com/tiagompalte/golang-clean-arch-template/internal/pkg/server/routes/v1"
	"github.com/tiagompalte/golang-clean-arch-template/pkg/server"
	"github.com/tiagompalte/golang-clean-arch-template/pkg/server/middleware"
)

func CreateRoute(app application.App) []server.GroupRoute {
	return []server.GroupRoute{
		{
			Path: "/api",
			Middlewares: []server.Middleware{
				middleware.LogMiddleware(app.Logger()),
			},
			GroupRoutes: []server.GroupRoute{
				v1.CreateRoute(app),
			},
		},
	}
}
