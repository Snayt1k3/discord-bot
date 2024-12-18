package roles

import (
	"log/slog"

	"github.com/bwmarrin/discordgo"
)

var reactionRoleMap = map[string]string{
	"🥒": "659728499195641857", // Роль: Овощ
}

var messageId = "1015669836891820043"  // cообщение за котором закреплены реакции


func OnMessageReactionAdd (s *discordgo.Session, r discordgo.MessageReactionAdd){
	if s.State.User.ID != r.UserID {
		return
	}

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

func OnMessageReactionRemove (s *discordgo.Session, r discordgo.MessageReactionRemove) {
	if s.State.User.ID != r.UserID {
		return
	}

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