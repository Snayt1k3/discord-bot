package config

import (
	"fmt"
	"os"
)

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
	TimeZone string
}

// Загрузка переменных из .env файла
func LoadConfig() (*DBConfig, error) {

	// Инициализируем конфигурацию из переменных окружения
	config := &DBConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
		TimeZone: os.Getenv("DB_TIMEZONE"),
	}

	// Проверяем, что все обязательные параметры заданы
	if config.Host == "" || config.Port == "" || config.User == "" || config.Password == "" || config.DBName == "" {
		return nil, fmt.Errorf("Missing required DB configuration")
	}

	return config, nil
}

// Формирование строки подключения (DSN) для PostgreSQL
func (c *DBConfig) DSN() string {
	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		c.Host, c.User, c.Password, c.DBName, c.Port, c.SSLMode, c.TimeZone,
	)
}