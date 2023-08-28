package v1

import (
	"github.com/tiagompalte/golang-clean-arch-template/application"
	"github.com/tiagompalte/golang-clean-arch-template/pkg/server"
)

func CreateRoute(app application.App) server.GroupRoute {
	return server.GroupRoute{
		Path: "/v1",
		GroupRoutes: []server.GroupRoute{
			CreateGroupTask(app),
			CreateGroupCategory(app),
		},
	}
}
