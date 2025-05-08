package handlers

import (
	"bot/internal/handlers/gachas/genshin"
	"bot/internal/handlers/gachas/wuwa"
	"bot/internal/handlers/guild"
	"bot/internal/interfaces"
	"github.com/bwmarrin/discordgo"
	"log/slog"
	"strings"
)

type CommandsDispatcher struct {
	guildKeeper interfaces.GuildKeeperInterface
	handlers    map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate) error

	guildHandlers   guild.GuildPreferencesHandlers
	genshinHandlers genshin.GenshinHandlers
	wuwaHandlers    wuwa.WuwaHandlers
}

func NewCommandsDispatcher(guildKeeper interfaces.GuildKeeperInterface, gachaAdapter interfaces.GachasAdapter) *CommandsDispatcher {
	return &CommandsDispatcher{
		guildKeeper:     guildKeeper,
		handlers:        map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate) error{},
		guildHandlers:   *guild.NewSettingsHandlers(guildKeeper),
		genshinHandlers: *genshin.NewGenshinHandlers(gachaAdapter),
		wuwaHandlers:    *wuwa.NewWuwaHandlers(gachaAdapter),
	}
}

func (cd *CommandsDispatcher) InitHandlers() {
	cd.handlers["help"] = HelpHandler
	cd.handlers["gachas"] = ShowSupportedGachas

	cd.guildHandlers.AddSettingsHandlers(cd.handlers)
	cd.genshinHandlers.AddGenshinHandlers(cd.handlers)
	cd.wuwaHandlers.AddWuwaHandlers(cd.handlers)
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
