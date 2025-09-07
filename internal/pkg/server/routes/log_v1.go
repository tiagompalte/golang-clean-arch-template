package routes

import (
	"net/http"

	"github.com/tiagompalte/golang-clean-arch-template/application"
	"github.com/tiagompalte/golang-clean-arch-template/internal/pkg/server/handler"
	"github.com/tiagompalte/golang-clean-arch-template/pkg/server"
)

func CreateGroupLogV1(app application.App) server.GroupRoute {
	routes := []server.Route{
		{
			Path:    "/",
			Method:  http.MethodPost,
			Handler: handler.CreateLogHandler(app.UseCase().CreateLogUseCase),
		},
	}

	return server.GroupRoute{
		Path:   "/logs",
		Routes: routes,
	}
}
