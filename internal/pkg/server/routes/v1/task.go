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
			Handler: handler.CreateTaskHandler(app.UseCase().CreateTask),
		},
		{
			Path:    "/",
			Method:  "GET",
			Handler: handler.FindAllTaskHandler(app.UseCase().FindAllTask),
		},
		{
			Path:    "/{uuid}",
			Method:  "GET",
			Handler: handler.FindOneTaskHandler(app.UseCase().FindOneTask),
		},
		{
			Path:    "/{uuid}/done",
			Method:  "PUT",
			Handler: handler.UpdateTaskDoneHandler(app.UseCase().UpdateTaskDone),
		},
		{
			Path:    "/{uuid}/undone",
			Method:  "PUT",
			Handler: handler.UpdateTaskUndoneHandler(app.UseCase().UpdateTaskUndone),
		},
		{
			Path:    "/{uuid}",
			Method:  "DELETE",
			Handler: handler.DeleteTaskHandler(app.UseCase().DeleteTask),
		},
	}

	return server.GroupRoute{
		Path:   "/tasks",
		Routes: routes,
	}
}
