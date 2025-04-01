package handlers

import (
	"bot/config"
	"bot/internal/adapters"
	"bot/internal/discord"
	dtoDiscord "bot/internal/dto/discord"
	"bot/internal/handlers/settings"
	"bot/internal/interfaces"
	"fmt"
	"log/slog"
	"math/rand"
	"time"

	"github.com/bwmarrin/discordgo"
)

type CommandsDispatcher struct {
	guildKeeper interfaces.GuildKeeperInterface
}

func NewCommandsDispatcher() *CommandsDispatcher {
	return &CommandsDispatcher{guildKeeper: adapters.NewServiceSettingsClient()}
}

func (cd *CommandsDispatcher) OnMemberJoin(s *discordgo.Session, u *discordgo.GuildMemberAdd) {
	settings, _ := cd.guildKeeper.GetGuildSettings(u.GuildID)

	randSource := rand.NewSource(time.Now().UnixNano())
	randGen := rand.New(randSource)

	randomIndex := randGen.Intn(len(config.HelloMessages))
	formattedMessage := fmt.Sprintf(config.HelloMessages[randomIndex], u.Member.Mention())

	discord.SendChannelMessage(settings.Settings.Welcome.ChannelId, formattedMessage)
}

func (cd *CommandsDispatcher) OnMessageReactionAdd(s *discordgo.Session, r *discordgo.MessageReactionAdd) {
	settings.OnMessageReactionAdd(cd.guildKeeper, s, r)
}

func (cd *CommandsDispatcher) OnMessageReactionRemove(s *discordgo.Session, r *discordgo.MessageReactionRemove) {
	settings.OnMessageReactionRemove(cd.guildKeeper, s, r)
}

func (cd *CommandsDispatcher) Dispatch(s *discordgo.Session, i *discordgo.InteractionCreate) {
	slog.Info("Handling interaction", "type", i.Type)

	data := dtoDiscord.HandlerData{Event: i, Session: s, Gk: cd.guildKeeper}

	switch i.Type {

	case discordgo.InteractionMessageComponent:
		if handler, ok := buttons[i.MessageComponentData().CustomID]; ok {
			handler(data)
		} else {
			slog.Warn("No handler found for button", "customID", i.MessageComponentData().CustomID)
		}

	default:
		cmdName := i.ApplicationCommandData().Name
		if handler, ok := adminHandlers[cmdName]; ok && cd.isAdmin(s, i.GuildID, i.Member.User.ID) {
			handler(data)
		} else if handler, ok := userHandlers[cmdName]; ok {
			if err := handler(data); err != nil {
				slog.Error("Error handling command", "err", err)
			}
		} else {
			slog.Warn("No handler found for command", "name", cmdName)
		}
	}
}

func (cd *CommandsDispatcher) isAdmin(s *discordgo.Session, guildID, userID string) bool {
	isAdmin, err := discord.IsAdmin(s, guildID, userID)
	if err != nil {
		slog.Error("Error checking admin status", "err", err)
		return false
	}
	return isAdmin
}
