package adapters

import (
	"context"
	"github.com/redis/go-redis/v9"
	"time"
)

// RedisAdapter impl interface RedisI.
type RedisAdapter struct {
	client *redis.Client
	ctx    context.Context
}

func NewRedisAdapter(addr, password string, db int) *RedisAdapter {
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})

	return &RedisAdapter{
		client: client,
		ctx:    context.Background(),
	}
}

func (r *RedisAdapter) Get(key string) (string, error) {
	return r.client.Get(r.ctx, key).Result()
}

func (r *RedisAdapter) Set(key string, value string, expiration time.Duration) error {
	return r.client.Set(r.ctx, key, value, expiration).Err()
}

func (r *RedisAdapter) Delete(key string) error {
	return r.client.Del(r.ctx, key).Err()
}

func (r *RedisAdapter) Exists(key string) (bool, error) {
	exists, err := r.client.Exists(r.ctx, key).Result()
	return exists > 0, err
}
