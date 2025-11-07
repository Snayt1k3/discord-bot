package discord

import (
	"bot/config"
	"log/slog"

	"github.com/bwmarrin/discordgo"
)

type DiscordBot struct {
	Session *discordgo.Session
}

var Bot *DiscordBot

func InitDiscordBot() {
	initBot()
	initConnection()
}

func initBot() {
	Bot = &DiscordBot{}

	session, err := discordgo.New("Bot " + config.GetDiscordToken())

	if err != nil {
		slog.Error("failed to create discord session", "error", err)
	}

	session.State.TrackVoice = true
	session.Identify.Intents = discordgo.IntentsAll
	Bot.Session = session

}

func initConnection() {
	if err := Bot.Session.Open(); err != nil {
		slog.Error("failed to create websocket connection to discord", "error", err)
		return
	}
}
