package handlers

import (
	"log/slog"
	"strings"

	"github.com/bwmarrin/discordgo"
)

type CommandsDispatcher struct {
	handlers map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate)
}

func NewCommandsDispatcher() *CommandsDispatcher {
	return &CommandsDispatcher{
		handlers: map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){},
	}
}

func (cd *CommandsDispatcher) InitHandlers(container *Container) {
	cd.handlers["help"] = container.Help
	cd.handlers["HelpPage"] = container.HelpPagination
	cd.handlers["HelpPageLast"] = container.HelpPaginationLast
	cd.handlers["HelpPageFirst"] = container.HelpPaginationFirst
	cd.handlers["menu"] = container.Menu
	cd.handlers["iprofile"] = container.InteractionProfile

	// Настройка Welcome
	cd.handlers["welcome-chnl"] = container.SetWelcomeChnl
	cd.handlers["welcomemsg-add"] = container.AddWelcomeMsg
	cd.handlers["welcomemsg-remove"] = container.RemoveWelcomeMsg

	// Настройка Roles/Reactions
	cd.handlers["rr-add"] = container.AddRole
	cd.handlers["rr-remove"] = container.RemoveRole
	cd.handlers["rr-message"] = container.SetRoleMsg

	// Logging
	cd.handlers["log-toggle"] = container.ToggleLog
	cd.handlers["log-chnl"] = container.AddLogChnl
	cd.handlers["log-chnl-rm"] = container.RemoveLogChnl

	cd.handlers["RolesReactionsSettings"] = container.RolesSettings
	cd.handlers["WelcomeSettings"] = container.WelcomePreferences
	cd.handlers["AutoModeSettings"] = container.ModerationSettings
	cd.handlers["LogSettings"] = container.LoggingSettings

	cd.handlers["automod-toggle"] = container.ToggleModeration
	cd.handlers["automod-bw-add"] = container.AddBannedWord
	cd.handlers["automod-bw-rm"] = container.RemoveBannedWord
	cd.handlers["automod-al-enable"] = container.AddAntiLinkChnl
	cd.handlers["automod-al-disable"] = container.RemoveAntiLinkChnl
	cd.handlers["automod-ac-enable"] = container.AddCapsLockChnl
	cd.handlers["automod-ac-disable"] = container.RemoveCapsLockChnl

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
		handler(s, i)
	} else {
		slog.Warn("No handler found for command", "name", name)
	}

}
