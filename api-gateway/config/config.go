package config

import (
	"log/slog"
	"os"
	"strconv"
)

type Config struct {
	GrpcHost  string
	GrpcPort  string
	RedisHost string
	RedisPort string
	RedisPass string
	RedisDB   int
}

func LoadConfig() (*Config, error) {
	redisDB, err := strconv.Atoi(os.Getenv("REDIS_DB"))

	if err != nil {
		slog.Warn("Failed to parse REDIS_DB", "error", err)
		redisDB = 0
	}

	config := &Config{
		GrpcHost:  os.Getenv("SETTINGS_ADDRESS"),
		GrpcPort:  os.Getenv("SETTINGS_PORT"),
		RedisHost: os.Getenv("REDIS_HOST"),
		RedisPort: os.Getenv("REDIS_PORT"),
		RedisPass: os.Getenv("REDIS_PASS"),
		RedisDB:   redisDB,
	}

	return config, nil
}
