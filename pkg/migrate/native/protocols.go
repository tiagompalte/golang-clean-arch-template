package nativemigrate

import (
	"context"
	"time"

	"github.com/tiagompalte/golang-clean-arch-template/pkg/repository"
)

type File interface {
	CreateUpAndDownFile(version int64, name string) error
	ListFileScripts() (map[string]Migrate, error)
	ReadScript(filename string) (string, error)
	PathFileUp(name string) string
	PathFileDown(name string) string
}

type RepositoryManager interface {
	Data() repository.DataSqlManager
	Version(conn repository.ConnectorSql) VersionRepository
}

type VersionRepository interface {
	CreateTable(ctx context.Context) error
	InsertBatch(ctx context.Context, names []string) error
	FindLastInserted(ctx context.Context) (string, error)
	IsAlreadyApply(ctx context.Context, name string) (bool, error)
	ExecScript(ctx context.Context, script string) error
	UpdateAppliedAt(ctx context.Context, name string, appliedAt *time.Time) error
}
