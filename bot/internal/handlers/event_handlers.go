package handlers

import (
	"bot/config"
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
	guildKeeper interfaces.GuildKeeperInterface
	cmds        []*discordgo.ApplicationCommand
}

func NewEventHandlers(guildKeeper interfaces.GuildKeeperInterface, cmds []*discordgo.ApplicationCommand) *EventHandlers {
	return &EventHandlers{guildKeeper: guildKeeper, cmds: cmds}
}

func (eh *EventHandlers) OnMessageReactionAdd(s *discordgo.Session, r *discordgo.MessageReactionAdd) {
	slog.Info("%v reacted with %v", r.UserID, r.Emoji.Name)

	guildSetting, err := eh.guildKeeper.GetGuildSettings(r.GuildID)

	if err != nil {
		slog.Error("Error while getting guild settings", "err", err)
		return
	}

	if r.MessageID != guildSetting.Settings.Roles.MessageId {
		return
	}

	roleId, exists := guildSetting.Settings.Roles.Matching[r.Emoji.ID]

	if !exists {
		return
	}

	err = s.GuildMemberRoleAdd(r.GuildID, r.UserID, roleId)

	if err != nil {
		slog.Error("Error adding role", "error", err)
	}
}

func (eh *EventHandlers) OnMessageReactionRemove(s *discordgo.Session, r *discordgo.MessageReactionRemove) {
	slog.Info("%v remove reaction %v", r.UserID, r.Emoji.Name)

	guildSetting, err := eh.guildKeeper.GetGuildSettings(r.GuildID)

	if err != nil {
		slog.Error("Error while getting guild settings", "err", err)
		return
	}

	if r.MessageID != guildSetting.Settings.Roles.MessageId {
		return
	}

	roleId, exists := guildSetting.Settings.Roles.Matching[r.Emoji.ID]

	if !exists {
		return
	}

	err = s.GuildMemberRoleRemove(r.GuildID, r.UserID, roleId)

	if err != nil {
		slog.Error("Error removing role", "error", err)
	}
}

func (eh *EventHandlers) OnMemberJoin(s *discordgo.Session, u *discordgo.GuildMemberAdd) {
	settings, _ := eh.guildKeeper.GetGuildSettings(u.GuildID)

	randSource := rand.NewSource(time.Now().UnixNano())
	randGen := rand.New(randSource)

	randomIndex := randGen.Intn(len(config.HelloMessages))
	formattedMessage := fmt.Sprintf(config.HelloMessages[randomIndex], u.Member.Mention())

	discord.SendChannelMessage(settings.Settings.Welcome.ChannelId, formattedMessage)
}

func (eh *EventHandlers) OnBotReady(s *discordgo.Session, g *discordgo.Ready) {
	err := s.UpdateCustomStatus(config.GetBotStatus())

	if err != nil {
		slog.Warn("failed to update custom status", "error", err)
	}

	for _, guild := range s.State.Guilds {
		slog.Info("Registering commands for server", "server_name", guild.Name, "server_id", guild.ID)
		for _, cmd := range eh.cmds {
			_, err := s.ApplicationCommandCreate(s.State.User.ID, guild.ID, cmd)
			if err != nil {
				slog.Error("Error while registering command", "command", cmd.Name, "server_name", guild.Name, "error", err)
			} else {
				slog.Info("Command registered successfully", "server_name", guild.Name, "command", cmd.Name)
			}
		}
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
