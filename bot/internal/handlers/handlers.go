package handlers

import (
	"log/slog"

	"github.com/bwmarrin/discordgo"
	"github.com/snayt1k3/discord-bot/bot/config"
)

const helpMessage string = ""

// ReadyHandler will be called when the bot receives the "ready" event from Discord.
func ReadyHandler(s *discordgo.Session, event *discordgo.Ready) {
	// Set the playing status.
	err := s.UpdateGameStatus(0, config.GetBotStatus())
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