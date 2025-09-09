package data

import (
	"github.com/tiagompalte/golang-clean-arch-template/internal/app/protocols"
)

type RepositoryManager interface {
	SqlManager | MongoManager
}

type Manager[T RepositoryManager] interface {
	Repository() T
}

type SqlManager struct {
	Task     func() protocols.TaskRepository
	Category func() protocols.CategoryRepository
	User     func() protocols.UserRepository
}

type MongoManager struct {
	Log func() protocols.LogRepository
}
