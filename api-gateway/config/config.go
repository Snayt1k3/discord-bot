package config

import (
	"os"
	"strconv"
)

type Config struct {
	GrpcHost string
	GrpcPort string
	RedisHost string
	RedisPort string
	RedisPass string
	RedisDB int
}


func LoadConfig() (*Config, error) {
	redisDB, err := strconv.Atoi(os.Getenv("REDIS_DB"))

	if err != nil {
		redisDB = 0 
	}


	config := &Config{
		GrpcHost: os.Getenv("GRPC_HOST"),
		GrpcPort: os.Getenv("GPRC_PORT"),
		RedisHost: os.Getenv("REDIS_HOST"),
		RedisPort: os.Getenv("REDIS_PORT"),
		RedisPass: os.Getenv("REDIS_PASS"),
		RedisDB: redisDB,
	}

	return config, nil
}