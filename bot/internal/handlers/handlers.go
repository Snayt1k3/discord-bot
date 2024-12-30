package handlers

import (
	"context"
	"log/slog"
	lv "bot/internal/lavalink"
	"bot/config"

	"github.com/bwmarrin/discordgo"
	"github.com/disgoorg/snowflake/v2"
)

const helpMessage string = ""

// ReadyHandler will be called when the bot receives the "ready" event from Discord.
func ReadyHandler(s *discordgo.Session, event *discordgo.Ready) {
	// Set the playing status.
	err := s.UpdateCustomStatus(config.GetBotStatus())
	if err != nil {
		slog.Warn("failed to update game status", "error", err)
	}
}

func MessageCreateHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID { // Preventing bot from using own commands
		return
	}

	slog.Info("processing command", "command", m.Content)

	// prefix := config.GetBotPrefix()
	// guildID := discord.SearchGuildByChannelID(m.ChannelID)

	// cmd := strings.Split(m.Content, " ") //	Splitting command into string slice
}

func OnVoiceStateUpdate(session *discordgo.Session, event *discordgo.VoiceStateUpdate) {
	if event.UserID != session.State.User.ID {
		return
	}

	var channelID *snowflake.ID
	if event.ChannelID != "" {
		id := snowflake.MustParse(event.ChannelID)
		channelID = &id
	}
	lv.LavalinkClient.Client.OnVoiceStateUpdate(context.TODO(), snowflake.MustParse(event.GuildID), channelID, event.SessionID)
	if event.ChannelID == "" {
		lv.LavalinkClient.Queue.Delete(event.GuildID)
	}
}

func OnVoiceServerUpdate(session *discordgo.Session, event *discordgo.VoiceServerUpdate) {
	lv.LavalinkClient.Client.OnVoiceServerUpdate(context.TODO(), snowflake.MustParse(event.GuildID), event.Token, event.Endpoint)
}