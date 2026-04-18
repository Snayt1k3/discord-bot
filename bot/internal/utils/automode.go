package utils

import (
	"bot/internal/dto"
	"fmt"
	"strings"
	"unicode"

	"github.com/bwmarrin/discordgo"
)

func AutomodeChecks(settings dto.AutoModeSettings, s *discordgo.Session, m *discordgo.MessageCreate) bool {
	if !settings.Enabled {
		return false
	}

	content := m.Content

	// 🚫 Проверка запрещённых слов
	for _, bw := range settings.BannedWords {
		if strings.Contains(strings.ToLower(content), strings.ToLower(bw.Word)) {
			s.ChannelMessageDelete(m.ChannelID, m.ID)
			SendTempMessage(s, m.ChannelID,
				fmt.Sprintf("❌ %s, your message contains a banned word!", m.Author.Mention()))
			return true
		}
	}

	// 🔗 Проверка ссылок
	for _, al := range settings.AntiLink {
		if strings.Contains(content, "http://") || strings.Contains(content, "https://") || strings.Contains(content, "discord.gg/") {
			if al.ChannelId == m.ChannelID {
				s.ChannelMessageDelete(m.ChannelID, m.ID)
				SendTempMessage(s, m.ChannelID,
					fmt.Sprintf("❌ %s, links are not allowed in this channel!", m.Author.Mention()),
				)
				return true
			}
		}
	}

	// 🔠 Проверка капслока
	for _, cl := range settings.CapsLock {
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
				SendTempMessage(s, m.ChannelID,
					fmt.Sprintf("❌ %s, please do not write in caps!", m.Author.Mention()),
				)
				return true
			}
		}
	}
	return false
}
