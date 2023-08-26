package v1

import (
	"github.com/tiagompalte/golang-clean-arch-template/application"
	"github.com/tiagompalte/golang-clean-arch-template/internal/pkg/server/handler"
	"github.com/tiagompalte/golang-clean-arch-template/pkg/server"
)

func CreateGroupTask(app application.App) server.GroupRoute {
	routes := make([]server.Route, 0)
	routes = append(routes, server.Route{
		Path:    "",
		Method:  "POST",
		Handler: handler.CreateTaskHandler(app.UseCase().CreateTask),
	})

	return server.GroupRoute{
		Path:   "/tasks",
		Routes: routes,
	}
}
