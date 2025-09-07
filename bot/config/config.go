package config

import (
	"fmt"
	"os"
)

type configuration struct {
	ApplicationId    string
	GuildId          string
	BotStatus        string
	DiscordToken     string
	ApiGatewayAddr   string
}

var config *configuration

func Load() {
	ApiGatewayAddr := os.Getenv("API_GATEWAY_ADDR")
	ApiGatewayPort := os.Getenv("API_GATEWAY_PORT")
	config = &configuration{
		ApplicationId:    os.Getenv("APPLICATION_ID"),
		DiscordToken:     os.Getenv("DISCORD_TOKEN"),
		GuildId:          os.Getenv("GUILD_ID"),
		BotStatus:        os.Getenv("BOT_STATUS"),
		ApiGatewayAddr:   fmt.Sprintf("%v:%v", ApiGatewayAddr, ApiGatewayPort),
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

func GetApiGatewayAddr() string {
	return config.ApiGatewayAddr
}
