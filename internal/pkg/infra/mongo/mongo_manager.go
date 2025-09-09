package mongo

import (
	"github.com/tiagompalte/golang-clean-arch-template/internal/app/protocols"
	"github.com/tiagompalte/golang-clean-arch-template/internal/pkg/infra/data"
	"github.com/tiagompalte/golang-clean-arch-template/pkg/repository"
)

type manager struct {
	log protocols.LogRepository
}

func NewRepositoryManager(conn repository.ConnectorMongo) data.Manager[data.MongoManager] {
	return manager{
		log: NewLogRepository(conn),
	}
}

func (m manager) Repository() data.MongoManager {
	return data.MongoManager{
		Log: func() protocols.LogRepository { return m.log },
	}
}
