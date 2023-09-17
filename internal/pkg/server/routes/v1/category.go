package v1

import (
	"github.com/tiagompalte/golang-clean-arch-template/application"
	"github.com/tiagompalte/golang-clean-arch-template/internal/pkg/server/handler"
	"github.com/tiagompalte/golang-clean-arch-template/pkg/server"
)

func CreateGroupCategory(app application.App) server.GroupRoute {
	routes := []server.Route{
		{
			Path:    "/",
			Method:  "GET",
			Handler: handler.FindAllCategoryHandler(app.UseCase().FindAllCategoryUseCase),
		},
	}

	return server.GroupRoute{
		Path:   "/categories",
		Routes: routes,
	}
}
