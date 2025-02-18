package data

import (
	"github.com/tiagompalte/golang-clean-arch-template/internal/app/protocols"
	"github.com/tiagompalte/golang-clean-arch-template/pkg/repository"
)

type RepositoryManager interface {
	Task() protocols.TaskRepository
	Category() protocols.CategoryRepository
	User() protocols.UserRepository
}

type repo struct {
	task     protocols.TaskRepository
	category protocols.CategoryRepository
	user     protocols.UserRepository
}

func NewRepositoryManager(conn repository.ConnectorSql) RepositoryManager {
	return repo{
		task:     NewTaskRepository(conn),
		category: NewCategoryRepository(conn),
		user:     NewUserRepository(conn),
	}
}

func (r repo) Task() protocols.TaskRepository {
	return r.task
}

func (r repo) Category() protocols.CategoryRepository {
	return r.category
}

func (r repo) User() protocols.UserRepository {
	return r.user
}
