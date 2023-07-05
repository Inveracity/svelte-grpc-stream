package cache

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type Cache struct {
	redis *redis.Client
}

func NewCache(redis *redis.Client) *Cache {

	return &Cache{redis: redis}
}

func (c *Cache) Set(key string, value string) error {
	ctx := context.Background()
	z := redis.ZAddArgs{
		Members: []redis.Z{
			{
				Score:  float64(time.Now().UnixNano()),
				Member: value,
			},
		},
	}

	ret := c.redis.ZAddArgs(ctx, key, z)

	if ret.Err() != nil {
		return ret.Err()
	}

	return nil
}

func (c *Cache) GetFrom(key, start, end string) ([]string, error) {
	ctx := context.Background()
	ret := c.redis.ZRangeByScore(ctx, key, &redis.ZRangeBy{
		Min: start,
		Max: end,
	})

	if ret.Err() != nil {
		return nil, ret.Err()
	}

	return ret.Val(), nil
}
