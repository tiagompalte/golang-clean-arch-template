package repository

import (
	"context"

	"github.com/tiagompalte/golang-clean-arch-template/internal/app/entity"
)

type CategoryRepository interface {
	Insert(ctx context.Context, category entity.Category) (uint32, error)
	FindBySlug(ctx context.Context, slug string) (entity.Category, error)
	FindAll(ctx context.Context) ([]entity.Category, error)
}
