package cache

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type Cache struct {
	ctx   context.Context
	redis *redis.Client
}

func NewCache(redis *redis.Client) *Cache {
	ctx := context.Background()
	return &Cache{ctx: ctx, redis: redis}
}

func (c *Cache) Set(key string, value string) error {
	z := redis.ZAddArgs{
		Members: []redis.Z{
			{
				Score:  float64(time.Now().UnixNano()),
				Member: value,
			},
		},
	}

	ret := c.redis.ZAddArgs(c.ctx, key, z)

	if ret.Err() != nil {
		return ret.Err()
	}

	return nil
}

func (c *Cache) GetFrom(key, start, end string) ([]string, error) {
	ret := c.redis.ZRangeByScore(c.ctx, key, &redis.ZRangeBy{
		Min:    start,
		Max:    end,
		Offset: 0,
		Count:  100,
	})

	if ret.Err() != nil {
		return nil, ret.Err()
	}

	return ret.Val(), nil
}
