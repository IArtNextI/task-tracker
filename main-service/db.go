package main

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

type Db interface {
	Exists(login string) (int64, error)
	SetNX(key string, value interface{}, expiration time.Duration) (bool, error)
}

type RedisDb struct{}

func (s *RedisDb) Exists(login string) (int64, error) {
	ctx := context.Background()
	rdb := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "e64a1e3065c49e781aa2721ce86e2725",
		DB:       0,
	})
	return rdb.Exists(ctx, login).Result()
}

func (s *RedisDb) SetNX(key string, value interface{}, expiration time.Duration) (bool, error) {
	ctx := context.Background()
	rdb := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "e64a1e3065c49e781aa2721ce86e2725",
		DB:       0,
	})
	return rdb.SetNX(ctx, key, value, expiration).Result()
}
