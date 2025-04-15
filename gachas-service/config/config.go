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
	TimeZone string
	GrpcPort string
}

func LoadConfig() (*Config, error) {
	config := &Config{
		Host:     os.Getenv("GACHA_POSTGRES_HOST"),
		Port:     os.Getenv("GACHA_POSTGRES_PORT"),
		User:     os.Getenv("GACHA_POSTGRES_USER"),
		Password: os.Getenv("GACHA_POSTGRES_PASSWORD"),
		DBName:   os.Getenv("GACHA_POSTGRES_DB"),
		SSLMode:  os.Getenv("GACHA_POSTGRES_SSLMODE"),
		TimeZone: os.Getenv("GACHA_POSTGRES_TIMEZONE"),
		GrpcPort: os.Getenv("GACHA_PORT"),
	}

	if config.Host == "" || config.Port == "" || config.User == "" || config.Password == "" || config.DBName == "" || config.GrpcPort == "" {
		return nil, fmt.Errorf("missing required configuration")
	}

	return config, nil
}

func (c *Config) DSN() string {
	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s timezone=%s",
		c.Host, c.User, c.Password, c.DBName, c.Port, c.SSLMode, c.TimeZone,
	)
}
