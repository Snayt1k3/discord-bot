package handlers

import (
	"bot/internal/interfaces"
	"bot/internal/utils"
	"github.com/bwmarrin/discordgo"
	"log/slog"
	"math/rand"
	"strings"
	"time"
)

// Хендлеры событий Discord отличных от Interaction событий
type EventHandlers struct {
	service interfaces.GuildServiceInterface
	cmds    []*discordgo.ApplicationCommand
}

func NewEventHandlers(service interfaces.GuildServiceInterface, cmds []*discordgo.ApplicationCommand) *EventHandlers {
	return &EventHandlers{service: service, cmds: cmds}
}

func (eh *EventHandlers) OnMessageReactionAdd(s *discordgo.Session, r *discordgo.MessageReactionAdd) {
	guildSetting, err := eh.service.GetGuildSettings(r.GuildID)

	if err != nil {
		slog.Error("Error while getting guild settings", "err", err)
		return
	}

	if r.MessageID != guildSetting.Roles.MessageID {
		return
	}

	roleId, exists := guildSetting.Roles.Matching[r.Emoji.ID]

	if !exists {
		return
	}

	err = s.GuildMemberRoleAdd(r.GuildID, r.UserID, roleId)

	if err != nil {
		slog.Error("Error adding role", "error", err)
	}
}

func (eh *EventHandlers) OnMessageReactionRemove(s *discordgo.Session, r *discordgo.MessageReactionRemove) {
	guildSetting, err := eh.service.GetGuildSettings(r.GuildID)

	if err != nil {
		slog.Error("Error while getting guild settings", "err", err)
		return
	}

	if r.MessageID != guildSetting.Roles.MessageID {
		return
	}

	roleId, exists := guildSetting.Roles.Matching[r.Emoji.ID]

	if !exists {
		return
	}

	err = s.GuildMemberRoleRemove(r.GuildID, r.UserID, roleId)

	if err != nil {
		slog.Error("Error removing role", "error", err)
	}
}

func (eh *EventHandlers) OnMemberJoin(s *discordgo.Session, u *discordgo.GuildMemberAdd) {
	settings, _ := eh.service.GetGuildSettings(u.GuildID)

	if len(settings.Welcome.Messages) == 0 {
		return
	}

	randSource := rand.NewSource(time.Now().UnixNano())
	randGen := rand.New(randSource)
	randomIndex := randGen.Intn(len(settings.Welcome.Messages))

	message := settings.Welcome.Messages[randomIndex]

	formattedMessage := strings.ReplaceAll(message, "{username}", u.User.Mention())

	utils.SendChannelMessage(settings.Welcome.ChannelID, formattedMessage)
}

func (eh *EventHandlers) OnGuildCreate(s *discordgo.Session, r *discordgo.GuildCreate) {
	err := eh.service.CreateSettings(r.ID)

	if err != nil {
		slog.Error("Error while creating guild settings", "err", err)
		return
	}
}
