package config

import (
	"fmt"
	"os"
	"strconv"
)

type configuration struct {
	ApplicationId    string
	GuildId          string
	BotStatus        string
	DiscordToken     string
	LavalinkAddr     string
	LavalinkPass     string
	LavalinkSecure   bool
	LavalinkNodeName string
	ApiGatewayAddr   string
}

var config *configuration

func Load() {
	LavalinkSecure, _ := strconv.ParseBool(os.Getenv("LAVALINK_SECURE"))
	ApiGatewayAddr := os.Getenv("API_GATEWAY_ADDR")
	ApiGatewayPort := os.Getenv("API_GATEWAY_PORT")
	config = &configuration{
		ApplicationId:    os.Getenv("APPLICATION_ID"),
		DiscordToken:     os.Getenv("DISCORD_TOKEN"),
		GuildId:          os.Getenv("GUILD_ID"),
		BotStatus:        os.Getenv("BOT_STATUS"),
		LavalinkAddr:     os.Getenv("LAVALINK_ADDR"),
		LavalinkPass:     os.Getenv("LAVALINK_PASS"),
		LavalinkNodeName: os.Getenv("LAVALINK_NODE_NAME"),
		LavalinkSecure:   LavalinkSecure,
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

func GetLavalinkAddr() string {
	return config.LavalinkAddr
}

func GetLavalinkPass() string {
	return config.LavalinkPass
}

func GetLavalinkNodeName() string {
	return config.LavalinkNodeName
}

func GetLavalinkSecure() bool {
	return config.LavalinkSecure
}

func GetApiGatewayAddr() string {
	return config.ApiGatewayAddr
}
