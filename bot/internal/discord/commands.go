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
}

// TODO: Добавить потом в main.go
func OnGuildCreate(s *discordgo.Session, g *discordgo.GuildCreate) {
	for _, command := range CommandsList {
		_, err := s.ApplicationCommandCreate(config.GetApplicationId(), g.Guild.ID, command)
		if err != nil {
			slog.Info("Error creating command '%s' for guild '%s': %s", command.Name, g.Guild.ID, err)
		} else {
			slog.Info("Command '%s' registered for new guild: %s", command.Name, g.Guild.ID)
		}
	}
}
