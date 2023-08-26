package repository

import (
	"context"

	"github.com/tiagompalte/golang-clean-arch-template/internal/app/entity"
)

type TaskRepository interface {
	Insert(ctx context.Context, task entity.Task) (uint32, error)
	FindByID(ctx context.Context, id uint32) (entity.Task, error)
	FindByUUID(ctx context.Context, uuid string) (entity.Task, error)
	FindAll(ctx context.Context) ([]entity.Task, error)
}
