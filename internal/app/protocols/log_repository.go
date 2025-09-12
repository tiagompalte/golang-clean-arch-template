package protocols

import (
	"context"

	"github.com/tiagompalte/golang-clean-arch-template/internal/app/entity"
)

type LogRepository interface {
	Insert(ctx context.Context, log entity.Log) (any, error)
	FindAll(ctx context.Context, limit int64) ([]entity.Log, error)
	FindByID(ctx context.Context, id string) (entity.Log, error)
}
