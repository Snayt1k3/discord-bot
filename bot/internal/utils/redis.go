package utils

import (
	"bot/config"
	"bot/internal/interfaces"
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

var ErrKeyNotFound = fmt.Errorf("redis: key not found")

type redisClient struct {
	rdb *redis.Client
}

// New creates a new Redis client and verifies the connection with PING.
func NewRedisClient() (interfaces.CacheClient, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     config.GetRedisAddr(),
		Password: config.GetRedisPassword(),
		DB:       config.GetRedisDB(),
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := rdb.Ping(ctx).Err(); err != nil {
		return nil, fmt.Errorf("redis: ping failed: %w", err)
	}

	return &redisClient{rdb: rdb}, nil
}

// Set stores value under key with an optional TTL (0 = persist).
func (c *redisClient) Set(ctx context.Context, key string, value any, ttl time.Duration) error {
	if err := c.rdb.Set(ctx, key, value, ttl).Err(); err != nil {
		return fmt.Errorf("redis Set %q: %w", key, err)
	}
	return nil
}

// Get retrieves the raw string value for key.
func (c *redisClient) Get(ctx context.Context, key string) (string, error) {
	val, err := c.rdb.Get(ctx, key).Result()
	if err == redis.Nil {
		return "", ErrKeyNotFound
	}
	if err != nil {
		return "", fmt.Errorf("redis Get %q: %w", key, err)
	}
	return val, nil
}

// SetJSON marshals value to JSON then stores it.
func (c *redisClient) SetJSON(ctx context.Context, key string, value any, ttl time.Duration) error {
	data, err := json.Marshal(value)
	if err != nil {
		return fmt.Errorf("redis SetJSON marshal %q: %w", key, err)
	}
	return c.Set(ctx, key, data, ttl)
}

// GetJSON retrieves a key and unmarshals the JSON payload into dest.
func (c *redisClient) GetJSON(ctx context.Context, key string, dest any) error {
	raw, err := c.Get(ctx, key)
	if err != nil {
		return err
	}
	if err := json.Unmarshal([]byte(raw), dest); err != nil {
		return fmt.Errorf("redis GetJSON unmarshal %q: %w", key, err)
	}
	return nil
}

// Delete removes one or more keys.
func (c *redisClient) Delete(ctx context.Context, keys ...string) (int64, error) {
	n, err := c.rdb.Del(ctx, keys...).Result()
	if err != nil {
		return 0, fmt.Errorf("redis Delete %v: %w", keys, err)
	}
	return n, nil
}

// Exists returns true if all provided keys exist.
func (c *redisClient) Exists(ctx context.Context, keys ...string) (bool, error) {
	n, err := c.rdb.Exists(ctx, keys...).Result()
	if err != nil {
		return false, fmt.Errorf("redis Exists %v: %w", keys, err)
	}
	return n == int64(len(keys)), nil
}

// Expire updates the TTL on an existing key.
func (c *redisClient) Expire(ctx context.Context, key string, ttl time.Duration) error {
	ok, err := c.rdb.Expire(ctx, key, ttl).Result()
	if err != nil {
		return fmt.Errorf("redis Expire %q: %w", key, err)
	}
	if !ok {
		return ErrKeyNotFound
	}
	return nil
}

// TTL returns the remaining lifetime of a key.
func (c *redisClient) TTL(ctx context.Context, key string) (time.Duration, error) {
	d, err := c.rdb.TTL(ctx, key).Result()
	if err != nil {
		return 0, fmt.Errorf("redis TTL %q: %w", key, err)
	}
	return d, nil
}

// Close closes the underlying connection pool.
func (c *redisClient) Close() error {
	return c.rdb.Close()
}
