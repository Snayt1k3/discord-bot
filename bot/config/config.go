package config

import (
	"log/slog"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type configuration struct {
	ApplicationId string
	GuildId string
	BotStatus string
	DiscordToken string
	LavalinkAddr string
	LavalinkPass string
	LavalinkSecure bool
	LavalinkNodeName string
}

var config *configuration

func Load() {
	err := godotenv.Load()
	// TODO: Убрать
	if err != nil {
		slog.Error("failed to load environment variables")
	}

	LavalinkSecure, _ := strconv.ParseBool(os.Getenv("LAVALINK_SECURE"))

	config = &configuration{
		ApplicationId: os.Getenv("APPLICATION_ID"),
		DiscordToken: os.Getenv("DISCORD_TOKEN"),
		GuildId: os.Getenv("GUILD_ID"),
		BotStatus: os.Getenv("BOT_STATUS"),
		LavalinkAddr: os.Getenv("LAVALINK_ADDR"),
		LavalinkPass: os.Getenv("LAVALINK_PASS"),
		LavalinkNodeName: os.Getenv("LAVALINK_NODE_NAME"),
		LavalinkSecure: LavalinkSecure,
	}
}

func GetGuildId() string {
	return config.GuildId
}

func GetApplicationId() string {
	return config.ApplicationId
}

func GetDiscordToken() string {
	return config.DiscordToken
}

func GetBotStatus() string {
	return config.BotStatus
}

func GetLavalinkAddr() string {
	return config.LavalinkAddr
}

func GetLavalinkPass()string {
	return config.LavalinkPass
}

func GetLavalinkNodeName()string {
	return config.LavalinkNodeName
}

func GetLavalinkSecure()bool {
	return config.LavalinkSecure
}