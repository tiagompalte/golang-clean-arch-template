package cache

import (
	"context"
	"sync"
	"time"
)

type item struct {
	value     interface{}
	createdAt int64
	ttl       int64
}

type MemoryCache struct {
	cache map[string]*item
	sync.RWMutex
}

// NewMemoryCache uses map to store key:value data in-memory.
func NewMemoryCache() *MemoryCache {
	c := &MemoryCache{cache: make(map[string]*item)}
	go c.setTtlTimer()

	return c
}

func (c *MemoryCache) setTtlTimer() {
	for {
		c.Lock()
		for k, v := range c.cache {
			if time.Now().Unix()-v.createdAt > v.ttl {
				delete(c.cache, k)
			}
		}
		c.Unlock()

		<-time.After(time.Second)
	}
}

func (c *MemoryCache) Set(ctx context.Context, key string, value any, ttl time.Duration) error {
	c.Lock()
	c.cache[key] = &item{
		value:     value,
		createdAt: time.Now().Unix(),
		ttl:       int64(ttl),
	}
	c.Unlock()

	return nil
}

func (c *MemoryCache) Get(ctx context.Context, key string) (any, error) {
	c.RLock()
	item, ex := c.cache[key]
	c.RUnlock()

	if !ex {
		return nil, ErrItemNotFound
	}

	return item.value, nil
}

func (c *MemoryCache) Clear(ctx context.Context, key string) error {
	c.Lock()
	c.cache[key] = nil
	c.Unlock()

	return nil
}

func (c *MemoryCache) ClearAll(ctx context.Context) error {
	c.Lock()
	c.cache = make(map[string]*item)
	c.Unlock()

	return nil
}

func (c *MemoryCache) IsHealthy(ctx context.Context) (bool, error) {
	return true, nil
}
