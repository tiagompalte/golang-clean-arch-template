package cache

import (
	"context"
	"time"
)

type MockCache struct {
}

func NewMockCache() Cache {
	return MockCache{}
}

func (c MockCache) Set(ctx context.Context, key string, value any, ttl time.Duration) error {
	return nil
}

func (c MockCache) Get(ctx context.Context, key string, value any) error {
	return ErrItemNotFound
}

func (c MockCache) Clear(ctx context.Context, key string) error {
	return nil
}

func (c MockCache) ClearAll(ctx context.Context) error {
	return nil
}

func (c MockCache) IsHealthy(ctx context.Context) (bool, error) {
	return true, nil
}
