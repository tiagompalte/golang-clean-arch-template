package sql

import (
	"github.com/tiagompalte/golang-clean-arch-template/internal/app/protocols"
	"github.com/tiagompalte/golang-clean-arch-template/internal/pkg/infra/data"
	"github.com/tiagompalte/golang-clean-arch-template/pkg/repository"
)

type manager struct {
	task     protocols.TaskRepository
	category protocols.CategoryRepository
	user     protocols.UserRepository
}

func NewRepositoryManager(conn repository.ConnectorSql) data.Manager[data.SqlManager] {
	return manager{
		task:     NewTaskRepository(conn),
		category: NewCategoryRepository(conn),
		user:     NewUserRepository(conn),
	}
}

func (m manager) Repository() data.SqlManager {
	return data.SqlManager{
		Task:     func() protocols.TaskRepository { return m.task },
		Category: func() protocols.CategoryRepository { return m.category },
		User:     func() protocols.UserRepository { return m.user },
	}
}
