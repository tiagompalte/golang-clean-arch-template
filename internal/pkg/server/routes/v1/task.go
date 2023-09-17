package v1

import (
	"github.com/tiagompalte/golang-clean-arch-template/application"
	"github.com/tiagompalte/golang-clean-arch-template/internal/pkg/server/handler"
	"github.com/tiagompalte/golang-clean-arch-template/pkg/server"
)

func CreateGroupTask(app application.App) server.GroupRoute {
	routes := []server.Route{
		{
			Path:    "/",
			Method:  "POST",
			Handler: handler.CreateTaskHandler(app.UseCase().CreateTaskUseCase),
		},
		{
			Path:    "/",
			Method:  "GET",
			Handler: handler.FindAllTaskHandler(app.UseCase().FindAllTaskUseCase),
		},
		{
			Path:    "/{uuid}",
			Method:  "GET",
			Handler: handler.FindOneTaskHandler(app.UseCase().FindOneTaskUseCase),
		},
		{
			Path:    "/{uuid}/done",
			Method:  "PUT",
			Handler: handler.UpdateTaskDoneHandler(app.UseCase().UpdateTaskDoneUseCase),
		},
		{
			Path:    "/{uuid}/undone",
			Method:  "PUT",
			Handler: handler.UpdateTaskUndoneHandler(app.UseCase().UpdateTaskUndoneUseCase),
		},
		{
			Path:    "/{uuid}",
			Method:  "DELETE",
			Handler: handler.DeleteTaskHandler(app.UseCase().DeleteTaskUseCase),
		},
	}

	return server.GroupRoute{
		Path:   "/tasks",
		Routes: routes,
	}
}
