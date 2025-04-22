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
		Host:     os.Getenv("GACHAS_POSTGRES_HOST"),
		Port:     os.Getenv("GACHAS_POSTGRES_PORT"),
		User:     os.Getenv("GACHAS_POSTGRES_USER"),
		Password: os.Getenv("GACHAS_POSTGRES_PASSWORD"),
		DBName:   os.Getenv("GACHAS_POSTGRES_DB"),
		SSLMode:  os.Getenv("GACHAS_POSTGRES_SSLMODE"),
		TimeZone: os.Getenv("GACHAS_POSTGRES_TIMEZONE"),
		GrpcPort: os.Getenv("GACHAS_PORT"),
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
