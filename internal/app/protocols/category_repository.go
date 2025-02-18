package protocols

import (
	"context"

	"github.com/tiagompalte/golang-clean-arch-template/internal/app/entity"
)

type CategoryRepository interface {
	Insert(ctx context.Context, category entity.Category) (uint32, error)
	FindBySlugAndUserID(ctx context.Context, slug string, userID uint32) (entity.Category, error)
	FindByUserID(ctx context.Context, userID uint32) ([]entity.Category, error)
	FindByID(ctx context.Context, id uint32) (entity.Category, error)
}
