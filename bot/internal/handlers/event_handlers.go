package handlers

import (
	"bot/internal/discord"
	"bot/internal/interfaces"
	"context"
	"fmt"
	"log/slog"
	"math/rand"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/disgoorg/snowflake/v2"
)

type EventHandlers struct {
	service interfaces.GuildServiceInterface
	cmds        []*discordgo.ApplicationCommand
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

	randSource := rand.NewSource(time.Now().UnixNano())
	randGen := rand.New(randSource)

	randomIndex := randGen.Intn(len(settings.Welcome.Messages))
	formattedMessage := fmt.Sprintf(settings.Welcome.Messages[randomIndex], u.Member.Mention())

	discord.SendChannelMessage(settings.Welcome.ChannelID, formattedMessage)
}

func (eh *EventHandlers) OnGuildCreate(s *discordgo.Session, r *discordgo.GuildCreate) {
	err := eh.service.CreateSettings(r.ID)
	
	if err != nil {
		slog.Error("Error while creating guild settings", "err", err)
		return
	}
}

func (eh *EventHandlers) OnVoiceStateUpdate(session *discordgo.Session, event *discordgo.VoiceStateUpdate) {
	if event.UserID != session.State.User.ID {
		return
	}

	var channelID *snowflake.ID

	if event.ChannelID != "" {
		id := snowflake.MustParse(event.ChannelID)
		channelID = &id
	}

	discord.Bot.Lavalink.OnVoiceStateUpdate(context.TODO(), snowflake.MustParse(event.GuildID), channelID, event.SessionID)

	if event.ChannelID == "" {
		discord.Bot.Queues.Delete(event.GuildID)
	}
}

func (eh *EventHandlers) OnVoiceServerUpdate(session *discordgo.Session, event *discordgo.VoiceServerUpdate) {
	discord.Bot.Lavalink.OnVoiceServerUpdate(context.TODO(), snowflake.MustParse(event.GuildID), event.Token, event.Endpoint)
}
