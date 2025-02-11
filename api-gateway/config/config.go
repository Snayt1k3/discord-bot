package config

import "os"

type Config struct {
	GrpcHost string
	GrpcPort string
}


func LoadConfig() (*Config, error) {
	config := &Config{
		GrpcHost: os.Getenv("GRPC_HOST"),
		GrpcPort: os.Getenv("GPRC_PORT"),
	}



	return config, nil
}