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
		Host:     os.Getenv("INTERACTION_POSTGRES_HOST"),
		Port:     os.Getenv("INTERACTION_POSTGRES_PORT"),
		User:     os.Getenv("INTERACTION_POSTGRES_USER"),
		Password: os.Getenv("INTERACTION_POSTGRES_PASSWORD"),
		DBName:   os.Getenv("INTERACTION_POSTGRES_DB"),
		GrpcPort: os.Getenv("INTERACTION_PORT"),
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
