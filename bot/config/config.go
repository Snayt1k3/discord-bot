package config

import (
	"log/slog"
	"os"

	"github.com/joho/godotenv"
)

type configuration struct {
	BotPrefix string
	BotStatus string
	DiscordToken string

}

var config *configuration

func Load() {
	err := godotenv.Load()

	if err != nil {
		slog.Error("failed to load environment variables")
	}

	config = &configuration{
		BotPrefix: os.Getenv("BOT_PREFIX"),
		DiscordToken: os.Getenv("DISCORD_TOKEN"),
		BotStatus: os.Getenv("BOT_STATUS"),
	}
}

func GetBotPrefix() string {
	return config.BotPrefix
}

func GetDiscordToken() string {
	return config.DiscordToken
}

func GetBotStatus() string {
	return config.BotStatus
}