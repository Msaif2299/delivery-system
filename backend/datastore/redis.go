package datastore

import (
	"context"
	"fmt"
	"os"

	"github.com/redis/go-redis/v9"
)

type RedisCache struct {
	conn *redis.Client
}

func NewRedisCache() *RedisCache {
	return &RedisCache{
		conn: redis.NewClient(
			&redis.Options{
				Addr:     fmt.Sprintf("redis:%s", os.Getenv("REDIS_PORT")),
				Password: os.Getenv("REDIS_PASSWORD"),
				DB:       0,
			},
		),
	}
}

func (c *RedisCache) Get(ctx context.Context, key string, val interface{}) error {
	err := c.conn.HGetAll(ctx, key).Scan(val)
	return err
}

func (c *RedisCache) Set(ctx context.Context, key string, value interface{}) error {
	_, err := c.conn.HSet(ctx, key, value).Result()
	return err
}
