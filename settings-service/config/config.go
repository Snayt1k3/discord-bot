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
	SSLMode  string

	GrpcPort string
}

func LoadConfig() (*Config, error) {
	config := &Config{
		Host:     os.Getenv("SETTINGS_POSTGRES_HOST"),
		Port:     os.Getenv("SETTINGS_POSTGRES_PORT"),
		User:     os.Getenv("SETTINGS_POSTGRES_USER"),
		Password: os.Getenv("SETTINGS_POSTGRES_PASSWORD"),
		DBName:   os.Getenv("SETTINGS_POSTGRES_DB"),
		SSLMode:  os.Getenv("SETTINGS_POSTGRES_SSLMODE"),

		GrpcPort: os.Getenv("SETTINGS_PORT"),
	}

	if config.Host == "" || config.Port == "" || config.User == "" || config.Password == "" || config.DBName == "" || config.GrpcPort == "" {
		return nil, fmt.Errorf("missing required configuration")
	}

	return config, nil
}

func (c *Config) DSN() string {
	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		c.Host, c.User, c.Password, c.DBName, c.Port, c.SSLMode,
	)
}
