package main

import (
	"context"
	"github.com/go-redis/redis/v8"
	"time"
)

type RedisClient struct {
	db  *redis.Client
	ctx context.Context
}

func CreateRedisClient(ctx context.Context, PoolSize int) *RedisClient {
	return &RedisClient{
		ctx: ctx,
		db: redis.NewClient(&redis.Options{
			Addr:     "localhost:6379",
			Password: "",
			DB:       0,
			PoolSize: PoolSize,
		}),
	}
}

func (redisClient *RedisClient) GetKey(key string) (string, error) {
	val, err := redisClient.db.Get(redisClient.ctx, key).Result()
	if err != nil || err == redis.Nil {
		return "", err
	}
	return val, err
}

func (redisClient *RedisClient) SetKey(key string, value interface{}, ttl time.Duration) (string, error) {
	status, err := redisClient.db.Set(redisClient.ctx, key, value, ttl).Result()
	if err != nil || err == redis.Nil {
		return "", err
	}
	return status, err
}
