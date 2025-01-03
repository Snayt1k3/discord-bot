package discord

import (
	"log/slog"

	"bot/config"

	"github.com/bwmarrin/discordgo"
)

var Session *discordgo.Session

func InitSession() {
	var err error
	Session, err = discordgo.New("Bot " + config.GetDiscordToken()) // Initializing discord session
	Session.State.TrackVoice = true
	Session.Identify.Intents = discordgo.IntentsAll
	if err != nil {
		slog.Error("failed to create discord session", "error", err)
	}

}

func InitConnection() {
	if err := Session.Open(); err != nil { // Creating a connection
		slog.Error("failed to create websocket connection to discord", "error", err)
		return
	}
}