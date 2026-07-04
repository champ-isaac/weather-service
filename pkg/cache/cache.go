package cache

import (
	"context"
	"sync"
	"time"
)

type Cache[T any] struct {
	m   sync.Map
	TTL time.Duration
}

type Item[T any] struct {
	Value     T
	ExpiredAt time.Time
}

func New[T any](ctx context.Context, ttl time.Duration) *Cache[T] {
	c := &Cache[T]{
		TTL: ttl,
	}

	go c.cleanUp(ctx)

	return c
}

func (c *Cache[T]) Set(key string, value T) {
	c.m.Store(key, Item[T]{
		Value:     value,
		ExpiredAt: time.Now().Add(c.TTL),
	})
}

func (c *Cache[T]) Get(key string) (T, bool) {
	var zero T
	v, ok := c.m.Load(key)
	if !ok {
		return zero, false
	}
	item := v.(Item[T])
	if time.Now().After(item.ExpiredAt) {
		return zero, false
	}
	return item.Value, true
}

func (c *Cache[T]) Delete(key string) {
	c.m.Delete(key)
}

func (c *Cache[T]) cleanUp(ctx context.Context) {
	ticket := time.NewTicker(c.TTL)
	defer ticket.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticket.C:
			now := time.Now()
			c.m.Range(func(k, v any) bool {
				item := v.(Item[T])
				if now.After(item.ExpiredAt) {
					c.m.Delete(k)
				}
				return true
			})
		}
	}
}
