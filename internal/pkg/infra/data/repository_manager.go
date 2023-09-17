package data

import (
	"github.com/tiagompalte/golang-clean-arch-template/internal/app/repository"
	pkgRepo "github.com/tiagompalte/golang-clean-arch-template/pkg/repository"
)

type RepositoryManager interface {
	Task() repository.TaskRepository
	Category() repository.CategoryRepository
	User() repository.UserRepository
}

type repo struct {
	task     repository.TaskRepository
	category repository.CategoryRepository
	user     repository.UserRepository
}

func NewRepositoryManager(conn pkgRepo.Connector) RepositoryManager {
	return repo{
		task:     NewTaskRepository(conn),
		category: NewCategoryRepository(conn),
		user:     NewUserRepository(conn),
	}
}

func (r repo) Task() repository.TaskRepository {
	return r.task
}

func (r repo) Category() repository.CategoryRepository {
	return r.category
}

func (r repo) User() repository.UserRepository {
	return r.user
}
