package nativemigrate

import "github.com/tiagompalte/golang-clean-arch-template/pkg/repository"

type repo struct {
	data repository.DataManager
}

func NewRepositoryManager(data repository.DataManager) RepositoryManager {
	return repo{data: data}
}

func (r repo) Data() repository.DataManager {
	return r.data
}

func (r repo) Version(conn repository.Connector) VersionRepository {
	return NewVersionRepositoryImpl(conn)
}
