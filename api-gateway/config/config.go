package config

import (
	"log/slog"
	"os"
	"strconv"
)

type Config struct {
	GrpcSettingsHost    string
	GrpcSettingsPort    string
	GrpcInteractionHost string
	GrpcInteractionPort string
	RedisHost           string
	RedisPort           string
	RedisPass           string
	RedisDB             int
	Port                string
}

func LoadConfig() (*Config, error) {
	redisDB, err := strconv.Atoi(os.Getenv("REDIS_DB"))

	if err != nil {
		slog.Warn("Failed to parse REDIS_DB", "error", err)
		redisDB = 0
	}

	config := &Config{
		GrpcSettingsHost:    os.Getenv("SETTINGS_ADDRESS"),
		GrpcSettingsPort:    os.Getenv("SETTINGS_PORT"),
		GrpcInteractionHost: os.Getenv("INTERACTION_ADDRESS"),
		GrpcInteractionPort: os.Getenv("INTERACTION_PORT"),
		RedisHost:           os.Getenv("REDIS_HOST"),
		RedisPort:           os.Getenv("REDIS_PORT"),
		RedisPass:           os.Getenv("REDIS_PASS"),
		RedisDB:             redisDB,
		Port:                os.Getenv("API_GATEWAY_PORT"),
	}

	return config, nil
}
