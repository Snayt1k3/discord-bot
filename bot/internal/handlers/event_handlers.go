package handlers

import (
	"bot/internal/discord"
	"bot/internal/dto"
	"bot/internal/http"
	"fmt"
	"log/slog"
	"math/rand"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"

	"bot/internal/utils"
)

type EventHandlers struct {
	http     *http.Container
	commands []*discordgo.ApplicationCommand
}

func NewEventHandlers(http *http.Container, commands []*discordgo.ApplicationCommand) *EventHandlers {
	return &EventHandlers{http: http, commands: commands}
}

func (eh *EventHandlers) OnMessageReactionAdd(s *discordgo.Session, r *discordgo.MessageReactionAdd) {
	guildSetting, err := eh.http.Settings.Get(r.GuildID)

	if err != nil {
		slog.Error("Error while getting http settings", "err", err)
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
	guildSetting, err := eh.http.Settings.Get(r.GuildID)

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
	settings, _ := eh.http.Settings.Get(u.GuildID)

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
	appID := s.State.User.ID

	for _, guild := range s.State.Guilds {
		for _, cmd := range discord.CommandsList {
			_, err := s.ApplicationCommandCreate(appID, guild.ID, cmd)
			if err != nil {
				slog.Error("Error creating command", "command", cmd.Name, "error", err)
			}
		}
	}
}

func (eh *EventHandlers) MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.Bot { return }

	res := eh.messageCheck(s, m)

	if !res {
		res, err := eh.http.Interaction.AddXP(m.GuildID, m.Author.ID, 10)
		
		if err != nil {
			slog.Error("Error while adding XP", "error", err)
			return
		}

		if res.LevelUp {
			utils.SendTempMessage(s, m.ChannelID,
				fmt.Sprintf("🎉 Congratulations %s! You've leveled up to level %d!", m.Author.Mention(), res.User.Level),
			)
		}
	}

}

func (eh *EventHandlers) MessageDelete(s *discordgo.Session, m *discordgo.MessageDelete) {
	fields := []*discordgo.MessageEmbedField{
		{Name: "Channel", Value: fmt.Sprintf("<#%s>", m.ChannelID), Inline: true},
		{Name: "Message ID", Value: m.ID, Inline: true},
	}

	if m.BeforeDelete != nil && m.BeforeDelete.Author != nil {
		fields = append(fields, &discordgo.MessageEmbedField{
			Name:   "Author",
			Value:  fmt.Sprintf("<@%s>", m.BeforeDelete.Author.ID),
			Inline: true,
		})
	}

	utils.SendLoggingMessage(eh.http, s, dto.MESSAGE_DELETE, m.GuildID,
		&discordgo.MessageEmbed{
			Title: "Message Deleted",
			Description: "A message was deleted.",
			Color: 0xFF0000,
			Fields: fields,
		},
	)
}

func (eh *EventHandlers) MessageUpdate(s *discordgo.Session, m *discordgo.MessageUpdate) {
	if m.BeforeUpdate == nil || m.BeforeUpdate.Content == m.Content {
		return
	}

	fields := []*discordgo.MessageEmbedField{
		{Name: "Channel", Value: fmt.Sprintf("<#%s>", m.ChannelID), Inline: true},
		{Name: "Message ID", Value: m.ID, Inline: true},
		{Name: "Before", Value: m.BeforeUpdate.Content, Inline: false},
		{Name: "After", Value: m.Content, Inline: false},
	}

	utils.SendLoggingMessage(eh.http, s, dto.MESSAGE_EDIT, m.GuildID,
		&discordgo.MessageEmbed{
			Title:       "Message Updated",
			Description: "A message was updated.",
			Color:       0x00FF00,
			Fields:      fields,
		},
	)
}

func (eh *EventHandlers) VoiceStatusChange(s *discordgo.Session, v *discordgo.VoiceStateUpdate) {
	beforeChannel := ""
	if v.BeforeUpdate != nil {
		beforeChannel = v.BeforeUpdate.ChannelID
	}

	afterChannel := v.ChannelID

	if beforeChannel == "" && afterChannel != "" {
		fields := []*discordgo.MessageEmbedField{
			{Name: "User", Value: fmt.Sprintf("<@%s>", v.UserID), Inline: true},
			{Name: "Channel", Value: fmt.Sprintf("<#%s>", afterChannel), Inline: true},
		}

		utils.SendLoggingMessage(eh.http, s, dto.JOIN_CHANNEL, v.GuildID,
			&discordgo.MessageEmbed{
				Title:       "Voice Join",
				Description: "User joined a voice channel.",
				Color:       0x00FF00,
				Fields:      fields,
			},
		)
		return
	}

	if beforeChannel != "" && afterChannel == "" {
		fields := []*discordgo.MessageEmbedField{
			{Name: "User", Value: fmt.Sprintf("<@%s>", v.UserID), Inline: true},
			{Name: "Channel", Value: fmt.Sprintf("<#%s>", beforeChannel), Inline: true},
		}

		utils.SendLoggingMessage(eh.http, s, dto.LEAVE_CHANNEL, v.GuildID,
			&discordgo.MessageEmbed{
				Title:       "Voice Leave",
				Description: "User left a voice channel.",
				Color:       0xFF0000,
				Fields:      fields,
			},
		)
		return
	}

	if beforeChannel != "" && afterChannel != "" && beforeChannel != afterChannel {
		fields := []*discordgo.MessageEmbedField{
			{Name: "User", Value: fmt.Sprintf("<@%s>", v.UserID), Inline: true},
			{Name: "From", Value: fmt.Sprintf("<#%s>", beforeChannel), Inline: true},
			{Name: "To", Value: fmt.Sprintf("<#%s>", afterChannel), Inline: true},
		}

		utils.SendLoggingMessage(eh.http, s, dto.MOVE_CHANNEL, v.GuildID,
			&discordgo.MessageEmbed{
				Title:       "Voice Move",
				Description: "User moved between voice channels.",
				Color:       0xFFFF00,
				Fields:      fields,
			},
		)
	}
}

func (eh *EventHandlers) OnInviteCreate(s *discordgo.Session, m *discordgo.InviteCreate) {
	fields := []*discordgo.MessageEmbedField{
		{Name: "Invite Link", Value: fmt.Sprintf("https://discord.gg/%s", m.Code)},
		{Name: "Created By", Value: fmt.Sprintf("<@%s>", m.Inviter.ID)},
	}

	utils.SendLoggingMessage(eh.http, s, dto.CREATE_INVITE, m.GuildID,
		&discordgo.MessageEmbed{
			Title:       "Invite Created",
			Description: "An invite was created.",
			Color:       0x00FF00,
			Fields:      fields,
		},
	)
}


func (eh *EventHandlers) GuildMemberRemove(s *discordgo.Session, m *discordgo.GuildMemberRemove) {
	fields := []*discordgo.MessageEmbedField{
		{Name: "User", Value: fmt.Sprintf("<@%s>", m.User.ID)},
	}

	utils.SendLoggingMessage(eh.http, s, dto.USER_LEAVE, m.GuildID, 
		&discordgo.MessageEmbed{
			Title:       "Member Left",
			Description: fmt.Sprintf("User %s left or was kicked from the server.", m.User.Username),
			Color:       0x808080,
			Fields:      fields,
		},
	)
}



func (eh *EventHandlers) messageCheck(s *discordgo.Session, m *discordgo.MessageCreate) bool {
	settings, err := eh.http.Settings.Get(m.GuildID)

	if err != nil {
		slog.Error("Error while fetching http settings", "err", err)
		return false
	}	
	
	return utils.AutomodeChecks(settings.AutoMode, s, m)
}
