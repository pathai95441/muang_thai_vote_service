package redis_cache

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

//go:generate mockgen -source=./redis_cache.go -destination=./mock/mock_redis_cache.go -package=mock_redis_cache
type IRedisCache interface {
	SetKeyValue(ctx context.Context, key string, value string) error
	GetValue(ctx context.Context, key string) (string, error)
}

type RedisCache struct {
	redis *redis.Client
}

func NewRedisCache(redis *redis.Client) RedisCache {
	return RedisCache{redis}
}

func (r RedisCache) SetKeyValue(ctx context.Context, key string, value string) error {

	err := r.redis.Set(ctx, key, value, 1*time.Hour).Err()
	if err != nil {
		return err
	}

	return nil
}

func (r RedisCache) GetValue(ctx context.Context, key string) (string, error) {

	value, err := r.redis.Get(ctx, key).Result()
	if err != nil {
		return "", err
	}

	return value, nil
}
