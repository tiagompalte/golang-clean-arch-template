package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/tiagompalte/golang-clean-arch-template/pkg/errors"
)

type RedisCache struct {
	redis  *redis.Client
	prefix string
}

func NewRedisCache(
	host string,
	port int,
	db int,
	pass string,
	prefix string,
) *RedisCache {
	return &RedisCache{
		prefix: prefix,
		redis: redis.NewClient(&redis.Options{
			Addr:     fmt.Sprintf("%s:%d", host, port),
			Password: pass,
			DB:       db,
		}),
	}
}

func (c *RedisCache) buildKey(key string) string {
	if c.prefix != "" {
		return c.prefix + ":" + key
	}

	return key
}

func (c *RedisCache) Set(ctx context.Context, key string, value any, ttl time.Duration) error {
	valueJson, err := json.Marshal(value)
	if err != nil {
		return errors.Wrap(err)
	}

	err = c.redis.Set(ctx, c.buildKey(key), valueJson, ttl).Err()
	if err != nil {
		return errors.Wrap(err)
	}

	return nil
}

func (c *RedisCache) Get(ctx context.Context, key string, value any) error {
	data, err := c.redis.Get(ctx, c.buildKey(key)).Bytes()
	if err == redis.Nil {
		return ErrItemNotFound
	}
	if err != nil {
		return errors.Wrap(err)
	}

	err = json.Unmarshal(data, &value)
	if err != nil {
		return errors.Wrap(err)
	}

	return nil
}

func (c *RedisCache) Clear(ctx context.Context, key string) error {
	err := c.redis.Del(ctx, key).Err()
	if err == redis.Nil {
		return ErrItemNotFound
	}
	if err != nil {
		return errors.Wrap(err)
	}

	return nil
}

func (c *RedisCache) ClearAll(ctx context.Context) error {
	err := c.redis.FlushAll(ctx).Err()
	if err == redis.Nil {
		return ErrItemNotFound
	}
	if err != nil {
		return errors.Wrap(err)
	}

	return nil
}

func (c *RedisCache) IsHealthy(ctx context.Context) (bool, error) {
	err := c.redis.Ping(ctx).Err()
	if err != nil {
		return false, errors.Wrap(err)
	}

	return true, nil
}
