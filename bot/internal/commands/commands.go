package commands

import (
	"bot/config"
	"bot/internal/handlers"
	"log/slog"

	"github.com/bwmarrin/discordgo"
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
		Name: 		 "skip",
		Description: "Skip current song",
	},

}


// TODO: Добавить потом в main.go
func OnGuildCreate(s *discordgo.Session, g *discordgo.GuildCreate) {
	for _, command := range CommandsList {
		_, err := s.ApplicationCommandCreate(config.GetApplicationId(), g.Guild.ID, command)
		if err != nil {
			slog.Info("Error creating command '%s' for guild '%s': %v", command.Name, g.Guild.ID, err)
		} else {
			slog.Info("Command '%s' registered for new guild: %s", command.Name, g.Guild.ID)
		}
	}
}


// CommandHandler маршрутизирует команды к соответствующим обработчикам
func CommandHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	switch i.ApplicationCommandData().Name {
	case "play":
		handlers.PlayCommandHandler(s, i)
	case "skip":
		handlers.SkipCommandHandler(s, i)
	case "stop":
		handlers.StopCommandHandler(s, i)
	}
}
