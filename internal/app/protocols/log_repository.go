package protocols

import (
	"context"

	"github.com/tiagompalte/golang-clean-arch-template/internal/app/entity"
)

type LogRepository interface {
	Insert(ctx context.Context, log entity.Log) (any, error)
	Find(ctx context.Context, filter any) ([]entity.Log, error)
}
