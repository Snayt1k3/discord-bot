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
	cd.handlers["toggle"] = container.ToggleFeature
	cd.handlers["menu"] = container.Menu

	// interactions
	cd.handlers["stats"] = container.UserStats
	cd.handlers["leaderboard"] = container.Leaderboard
	cd.handlers["LeaderboardPage"] = container.LeaderboardPagination
	cd.handlers["LeaderboardLast"] = container.LeaderboardPaginationLast
	cd.handlers["LeaderboardFirst"] = container.LeaderboardPaginationFirst

	// Welcome
	cd.handlers["welcome-chnl"] = container.SetWelcomeChnl
	cd.handlers["welcome-msg"] = container.WelcomeMsg

	// Roles/Reactions
	cd.handlers["rr-add"] = container.AddRole
	cd.handlers["rr-remove"] = container.RemoveRole
	cd.handlers["rr-message"] = container.SetRoleMsg

	// Logging
	cd.handlers["log-edit"] = container.LogEdit

	// Moderation
	cd.handlers["automod-bannedword"] = container.BannedWord
	cd.handlers["automod-antilink"] = container.AntiLink
	cd.handlers["automod-anticaps"] = container.AntiCaps

	// Info
	cd.handlers["user-info"] = container.UserInfo
	cd.handlers["server-info"] = container.ServerInfo

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
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
        Type: discordgo.InteractionResponseChannelMessageWithSource,
        Data: &discordgo.InteractionResponseData{
            Content: "Command not found.",
            Flags:   discordgo.MessageFlagsEphemeral,
        },
    })
	}

}
