package discord

import (
	"bot/config"
	"github.com/bwmarrin/discordgo"
	"log/slog"
)

var CommandsList = []*discordgo.ApplicationCommand{
	{
		Name:        "play",
		Description: "Play a song or add it to the queue",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "query",
				Description: "The name or URL of the song",
				Required:    true,
			},
		},
	},
	{
		Name:        "stop",
		Description: "Stop playing music",
	},
	{
		Name:        "skip",
		Description: "Skip the current song",
	},
	{
		Name:        "help",
		Description: "Display a list of available commands",
	},
	{
		Name:        "resume",
		Description: "Resume the paused song",
	},
	{
		Name:        "pause",
		Description: "Pause the currently playing song",
	},
	{
		Name:        "settings",
		Description: "Configure server settings. Only for admins",
	},
}

func OnGuildCreate(s *discordgo.Session, g *discordgo.GuildCreate) {
	for _, command := range CommandsList {
		_, err := s.ApplicationCommandCreate(config.GetApplicationId(), g.Guild.ID, command)
		if err != nil {
			slog.Error("Error creating command for guild", "command", command.Name, "guild_id", g.Guild.ID, "error", err)
		} else {
			slog.Info("Command registered for new guild", "command", command.Name, "guild_id", g.Guild.ID)
		}
	}
}
