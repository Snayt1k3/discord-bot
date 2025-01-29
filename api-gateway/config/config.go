package config

import "os"

type Config struct {
	SettingsAddress string
	SettingsPort string
}


func LoadConfig() (*Config, error) {
	config := &Config{
		SettingsAddress: os.Getenv("SETTINGS_ADDRESS"),
		SettingsPort: os.Getenv("SETTINGS_PORT"),
	}



	return config, nil
}