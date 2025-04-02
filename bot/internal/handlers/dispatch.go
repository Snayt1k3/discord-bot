package handlers

import (
	dtoDiscord "bot/internal/dto/discord"
	"bot/internal/interfaces"
	"bot/internal/handlers/gachas"
	"bot/internal/handlers/settings"
	"log/slog"
	"github.com/bwmarrin/discordgo"
)

type CommandsDispatcher struct {
	guildKeeper interfaces.GuildKeeperInterface
	handlers map[string]func(data dtoDiscord.HandlerData) error
}

func NewCommandsDispatcher(gk interfaces.GuildKeeperInterface) *CommandsDispatcher {
	return &CommandsDispatcher{guildKeeper: gk, handlers: map[string]func(data dtoDiscord.HandlerData) error{}}
}

func (cd *CommandsDispatcher) InitHandlers() {
	cd.handlers["help"] = HelpHandler
	
	gachas.AddHandlers(cd.handlers)
	settings.AddSettingsHandlers(cd.handlers)
}

func (cd *CommandsDispatcher) Dispatch(s *discordgo.Session, i *discordgo.InteractionCreate) {
	slog.Info("Startuing handle interaction", "type", i.Type)
	
	var name string
	data := dtoDiscord.HandlerData{Event: i, Session: s, Gk: cd.guildKeeper}

	switch i.Type {

	case discordgo.InteractionMessageComponent:
		name = i.MessageComponentData().CustomID
	
	default:
		name = i.ApplicationCommandData().Name
	}

	if handler, ok := cd.handlers[name]; ok {
	
		if err := handler(data); err != nil {
			slog.Error("Error handling command", "err", err, "name", name)
		} 

	} else {
		slog.Warn("No handler found for command", "name", name)
	}

}

