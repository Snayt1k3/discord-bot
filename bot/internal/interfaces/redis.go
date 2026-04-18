package interfaces

import (
	"context"
	"time"
)

// CacheClient defines the interface for Redis CRUD operations.
type CacheClient interface {
	// Set stores a value under the given key with an optional TTL (0 = no expiry).
	Set(ctx context.Context, key string, value any, ttl time.Duration) error

	// Get retrieves the raw string value for a key.
	// Returns ErrKeyNotFound if the key does not exist.
	Get(ctx context.Context, key string) (string, error)

	// GetJSON retrieves a JSON-encoded value and unmarshals it into dest.
	GetJSON(ctx context.Context, key string, dest any) error

	// SetJSON marshals value to JSON and stores it under key with an optional TTL.
	SetJSON(ctx context.Context, key string, value any, ttl time.Duration) error

	// Delete removes one or more keys. Returns the number of keys deleted.
	Delete(ctx context.Context, keys ...string) (int64, error)

	// Exists reports whether all the given keys exist.
	Exists(ctx context.Context, keys ...string) (bool, error)

	// Expire updates the TTL of an existing key.
	Expire(ctx context.Context, key string, ttl time.Duration) error

	// TTL returns the remaining time-to-live of a key.
	// Returns -1 if the key has no expiry, -2 if the key does not exist.
	TTL(ctx context.Context, key string) (time.Duration, error)

	// Close releases the underlying connection.
	Close() error
}
