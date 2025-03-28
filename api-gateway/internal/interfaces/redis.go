package interfaces

import "time"

type RedisInterface interface {
	Get(key string) (string, error)
	Set(key string, value string, expiration time.Duration) error
	Delete(key string) error
	Exists(key string) (bool, error)
}
