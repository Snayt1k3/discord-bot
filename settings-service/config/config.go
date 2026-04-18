package config

import (
	"fmt"
	"os"
)

type Config struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string

	GrpcPort string
}

func LoadConfig() (*Config, error) {
	config := &Config{
		Host:     os.Getenv("POSTGRES_HOST"),
		Port:     os.Getenv("POSTGRES_PORT"),
		User:     os.Getenv("POSTGRES_USER"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		DBName:   os.Getenv("SETTINGS_POSTGRES_DB"),

		GrpcPort: os.Getenv("SETTINGS_GRPC_PORT"),
	}

	if config.Host == "" || config.Port == "" || config.User == "" || config.Password == "" || config.DBName == "" || config.GrpcPort == "" {
		return nil, fmt.Errorf("missing required configuration")
	}

	return config, nil
}

func (c *Config) DSN() string {
	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s",
		c.Host, c.User, c.Password, c.DBName, c.Port,
	)
}
