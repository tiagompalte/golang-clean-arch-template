package cache

import (
	"context"
	"time"
)

type Cache interface {
	Set(ctx context.Context, key string, value any, ttl time.Duration) error
	Get(ctx context.Context, key string) (any, error)
	Clear(ctx context.Context, key string) error
	ClearAll(ctx context.Context) error
}
