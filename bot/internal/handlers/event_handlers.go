package handlers

import (
	"bot/internal/discord"
	"bot/internal/http"
	"fmt"
	"log/slog"
	"math/rand"
	"strings"
	"time"
	"unicode"

	"github.com/bwmarrin/discordgo"

	"bot/internal/utils"
)

// EventHandlers Хендлеры событий Discord отличных от Interaction событий
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
	if m.Author.Bot {
		return
	}

	res := eh.automodeCheck(s, m)

	// Adding XP to user
	if !res {
		_, err := eh.http.Interaction.AddXP(m.GuildID, m.Author.ID, 10)
		if err != nil {
			slog.Error("Error while adding XP", "error", err)
		}
	}

}

func (eh *EventHandlers) MessageDelete(s *discordgo.Session, m *discordgo.MessageDelete) {
	fields := []*discordgo.MessageEmbedField{
		{Name: "Channel", Value: fmt.Sprintf("<#%s>", m.ChannelID), Inline: true},
		{Name: "Message ID", Value: m.ID, Inline: true},
	}

	// Try to include the author if cached
	if m.BeforeDelete != nil && m.BeforeDelete.Author != nil {
		fields = append(fields, &discordgo.MessageEmbedField{
			Name:   "Author",
			Value:  fmt.Sprintf("<@%s>", m.BeforeDelete.Author.ID),
			Inline: true,
		})
	}

	_ = eh.sendLogMessage(s, m.GuildID,
		"Message Deleted",
		"A message was deleted.",
		0xFF0000, // red
		fields,
	)
}

func (eh *EventHandlers) MessageDeleteBulk(s *discordgo.Session, m *discordgo.MessageDeleteBulk) {
	fields := []*discordgo.MessageEmbedField{
		{Name: "Channel", Value: fmt.Sprintf("<#%s>", m.ChannelID), Inline: true},
		{Name: "Messages Count", Value: fmt.Sprintf("%d", len(m.Messages)), Inline: true},
	}
	_ = eh.sendLogMessage(s, m.GuildID,
		"Bulk Message Deletion",
		"Multiple messages were deleted.",
		0xFF4500, // orange
		fields,
	)
}

func (eh *EventHandlers) OnInviteCreate(s *discordgo.Session, m *discordgo.InviteCreate) {
	fields := []*discordgo.MessageEmbedField{
		{Name: "Invite Link", Value: fmt.Sprintf("https://discord.gg/%s", m.Code)},
		{Name: "Created By", Value: fmt.Sprintf("<@%s>", m.Inviter.ID)},
	}
	_ = eh.sendLogMessage(s, m.GuildID,
		"Invite Created",
		fmt.Sprintf("An invite was created for channel <#%s>.", m.ChannelID),
		0x00FF00, // green
		fields,
	)
}

func (eh *EventHandlers) GuildBanAdd(s *discordgo.Session, m *discordgo.GuildBanAdd) {
	fields := []*discordgo.MessageEmbedField{
		{Name: "User", Value: fmt.Sprintf("<@%s>", m.User.ID)},
	}
	_ = eh.sendLogMessage(s, m.GuildID,
		"User Banned",
		fmt.Sprintf("User %s was banned from the server.", m.User.Username),
		0x8B0000, // dark red
		fields,
	)
}

func (eh *EventHandlers) GuildBanRemove(s *discordgo.Session, m *discordgo.GuildBanRemove) {
	fields := []*discordgo.MessageEmbedField{
		{Name: "User", Value: fmt.Sprintf("<@%s>", m.User.ID)},
	}
	_ = eh.sendLogMessage(s, m.GuildID,
		"Ban Removed",
		fmt.Sprintf("The ban for user %s has been lifted.", m.User.Username),
		0x00CED1, // teal
		fields,
	)
}

func (eh *EventHandlers) GuildMemberRemove(s *discordgo.Session, m *discordgo.GuildMemberRemove) {
	fields := []*discordgo.MessageEmbedField{
		{Name: "User", Value: fmt.Sprintf("<@%s>", m.User.ID)},
	}
	_ = eh.sendLogMessage(s, m.GuildID,
		"Member Left",
		fmt.Sprintf("User %s left or was kicked from the server.", m.User.Username),
		0x808080, // gray
		fields,
	)
}

// Вспомогательная функция для отправки логов
func (eh *EventHandlers) sendLogMessage(
	s *discordgo.Session,
	guildId string,
	title string,
	description string,
	color int,
	fields []*discordgo.MessageEmbedField,
) error {

	settings, err := eh.http.Settings.Get(guildId)

	if err != nil {
		slog.Error("Error while getting settings", "error", err)
		return err
	}

	if !settings.Log.Enabled {
		return nil
	}

	embed := &discordgo.MessageEmbed{
		Title:       title,
		Description: description,
		Color:       color,
		Timestamp:   time.Now().Format(time.RFC3339),
		Footer: &discordgo.MessageEmbedFooter{
			Text: "Logs Bots",
		},
		Fields: fields,
	}

	_, err = s.ChannelMessageSendEmbed(settings.Log.ChannelID, embed)

	if err != nil {
		slog.Error("Error while sending log message", "error", err)
	}

	return nil
}

// Вспомогательная функция для проверки всех условий автомодерации.
func (eh *EventHandlers) automodeCheck(s *discordgo.Session, m *discordgo.MessageCreate) bool {
	settings, err := eh.http.Settings.Get(m.GuildID)

	if err != nil {
		slog.Error("Error while fetching http settings", "err", err)
		return false
	}

	autoMode := settings.AutoMode
	if !autoMode.Enabled {
		return false
	}

	content := m.Content

	// 🚫 Проверка запрещённых слов
	for _, bw := range autoMode.BannedWords {
		if strings.Contains(strings.ToLower(content), strings.ToLower(bw.Word)) {
			s.ChannelMessageDelete(m.ChannelID, m.ID)
			utils.SendTempMessage(s, m.ChannelID,
				fmt.Sprintf("❌ %s, your message contains a banned word!", m.Author.Mention()))
			return true
		}
	}

	// 🔗 Проверка ссылок
	for _, al := range autoMode.AntiLink {
		if strings.Contains(content, "http://") || strings.Contains(content, "https://") || strings.Contains(content, "discord.gg/") {
			if al.ChannelId == m.ChannelID {
				s.ChannelMessageDelete(m.ChannelID, m.ID)
				utils.SendTempMessage(s, m.ChannelID,
					fmt.Sprintf("❌ %s, links are not allowed in this channel!", m.Author.Mention()))
				return true
			}
		}
	}

	// 🔠 Проверка капслока
	for _, cl := range autoMode.CapsLock {
		if cl.ChannelId == m.ChannelID {
			upperCount := 0
			letters := 0
			for _, r := range content {
				if unicode.IsLetter(r) {
					letters++
					if unicode.IsUpper(r) {
						upperCount++
					}
				}
			}
			if letters > 0 && float64(upperCount)/float64(letters) > 0.7 {
				s.ChannelMessageDelete(m.ChannelID, m.ID)
				utils.SendTempMessage(s, m.ChannelID,
					fmt.Sprintf("❌ %s, please do not write in caps!", m.Author.Mention()))
				return true
			}
		}
	}

	return false
}
