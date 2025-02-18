package protocols

import (
	"context"

	"github.com/tiagompalte/golang-clean-arch-template/internal/app/entity"
)

type UserRepository interface {
	Insert(ctx context.Context, user entity.User, passEncrypted string) (uint32, error)
	FindByID(ctx context.Context, id uint32) (entity.User, error)
	FindByUUID(ctx context.Context, uuid string) (entity.User, error)
	FindByEmail(ctx context.Context, email string) (entity.User, error)
	GetPassEncryptedByEmail(ctx context.Context, email string) (string, error)
	UpdateName(ctx context.Context, user entity.User) error
}
