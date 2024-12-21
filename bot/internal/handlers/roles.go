package handlers

import (

	"log/slog"

	"github.com/bwmarrin/discordgo"
)

var reactionRoleMap = map[string]string{
	"🥒": "659728499195641857", // Роль: Овощ
	"valorant_logo": "", // Роль: Валорантер
	"dota2": "", // Роль: Дотер
	"rainbow_smash": "", // Роль: Фортнайтер
	"sanhua_angry": "", // Роль: Вувщик
	"Sucrose": "", // Роль: Геншиненок
	"pompomgallerythenwaketoweep4": "1319978883386441779", // Роль: Хсрщик
}

var messageId = "1015669836891820043" // cообщение за котором закреплены реакции

func OnMessageReactionAdd(s *discordgo.Session, r *discordgo.MessageReactionAdd) {
	slog.Info("%v reacted with %v", r.UserID, r.Emoji.Name)

	if r.MessageID != messageId {
		return
	}

	roleId, exists := reactionRoleMap[r.Emoji.Name]

	if !exists {
		return
	}

	err := s.GuildMemberRoleAdd(r.GuildID, r.UserID, roleId)

	if err != nil {
		slog.Warn("Error adding role", "error", err)
	}

}

func OnMessageReactionRemove(s *discordgo.Session, r *discordgo.MessageReactionRemove) {
	slog.Info("%v remove reaction %v", r.UserID, r.Emoji.Name)

	if r.MessageID != messageId {
		return
	}

	roleId, exists := reactionRoleMap[r.Emoji.Name]

	if !exists {
		return
	}

	err := s.GuildMemberRoleRemove(r.GuildID, r.UserID, roleId)

	if err != nil {
		slog.Warn("Error removing role", "error", err)
	}
}
