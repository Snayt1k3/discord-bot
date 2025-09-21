package handlers

import (
	"log/slog"
	"strings"

	"github.com/bwmarrin/discordgo"

	"bot/internal/handlers/guild"
	guildAdapters "bot/internal/adapters/guild"
)

type CommandsDispatcher struct {
	guildAdapter guildAdapters.GuildAdapter
	handlers     map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate) error // todo: Change naming
	Handlers guild.Handlers
}

func NewCommandsDispatcher(guildService guildAdapters.GuildAdapter) *CommandsDispatcher {
	return &CommandsDispatcher{
		guildAdapter: guildService,
		handlers:     map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate) error{},
		Handlers:     *guild.NewHandlers(guildService),
	}
}

func (cd *CommandsDispatcher) AddHandler(name string, handler func(s *discordgo.Session, i *discordgo.InteractionCreate) error) {
	cd.handlers[name] = handler
}

func (cd *CommandsDispatcher) InitHandlers() {
	cd.AddHandler("help", HelpHandler)
	cd.Handlers.AddHandlers(cd.handlers)
}

func (cd *CommandsDispatcher) Dispatch(s *discordgo.Session, i *discordgo.InteractionCreate) {
	slog.Info("Starting handle interaction", "type", i.Type)

	var name string

	switch i.Type {

	case discordgo.InteractionMessageComponent:
		name, _, _ = strings.Cut(i.MessageComponentData().CustomID, "_")
		// logic Name_1: after underscore it's an id or some useful info
		// only for buttons

	default:
		name = i.ApplicationCommandData().Name
	}

	slog.Info("Cmd:", "name", name)

	if handler, ok := cd.handlers[name]; ok {

		if err := handler(s, i); err != nil {
			slog.Error("Error handling command", "err", err, "name", name)
		}

	} else {
		slog.Warn("No handler found for command", "name", name)
	}

}
