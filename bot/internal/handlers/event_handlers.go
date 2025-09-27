package handlers

import (
	"fmt"
	"log/slog"
	"math/rand"
	"strings"
	"time"
	"unicode"

	"github.com/bwmarrin/discordgo"

	"bot/internal/adapters/guild"
	"bot/internal/utils"
)

// Хендлеры событий Discord отличных от Interaction событий
type EventHandlers struct {
	service guild.GuildAdapter
	cmds    []*discordgo.ApplicationCommand
}

func NewEventHandlers(service guild.GuildAdapter, cmds []*discordgo.ApplicationCommand) *EventHandlers {
	return &EventHandlers{service: service, cmds: cmds}
}

func (eh *EventHandlers) OnMessageReactionAdd(s *discordgo.Session, r *discordgo.MessageReactionAdd) {
	guildSetting, err := eh.service.Settings.Get(r.GuildID)

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
	guildSetting, err := eh.service.Settings.Get(r.GuildID)

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
	settings, _ := eh.service.Settings.Get(u.GuildID)

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
	err := eh.service.Settings.Create(r.ID)

	if err != nil {
		slog.Error("Error while creating guild settings", "err", err)
		return
	}
}

func (eh *EventHandlers) MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.Bot {
		return
	}

	eh.automodeCheck(s, m)

}

// Вспомогательная функция для проверки всех условий автомодерации.
func (eh *EventHandlers) automodeCheck(s *discordgo.Session, m *discordgo.MessageCreate) {
	settings, err := eh.service.Settings.Get(m.GuildID)

	if err != nil {
		slog.Error("Error while fetching guild settings", "err", err)
		return
	}

	autoMode := settings.AutoMode
	if !autoMode.Enabled {
		return
	}

	content := m.Content

	for _, bw := range autoMode.BannedWords {
		if strings.Contains(strings.ToLower(content), strings.ToLower(bw.Word)) {
			s.ChannelMessageDelete(m.ChannelID, m.ID)
			utils.SendTempMessage(s, m.ChannelID, fmt.Sprintf("❌ %s, your message contains a banned word!", m.Author.Mention()))
			return
		}
	}

	for _, al := range autoMode.AntiLink {
		if strings.Contains(content, "http://") || strings.Contains(content, "https://") || strings.Contains(content, "discord.gg/") {
			if al.ChannelId == m.ChannelID {
				s.ChannelMessageDelete(m.ChannelID, m.ID)
				utils.SendTempMessage(s, m.ChannelID, fmt.Sprintf("❌ %s, links are not allowed in this channel!", m.Author.Mention()))
				return
			}
		}
	}

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
				utils.SendTempMessage(s, m.ChannelID, fmt.Sprintf("❌ %s, please do not write in caps!", m.Author.Mention()))
				return
			}
		}
	}
}
