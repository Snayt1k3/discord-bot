package handlers

import (
	"bot/config"
	"github.com/bwmarrin/discordgo"
	"log/slog"
)

func OnMessageReactionAdd(s *discordgo.Session, r *discordgo.MessageReactionAdd) {
	slog.Info("%v reacted with %v", r.UserID, r.Emoji.Name)

	if r.MessageID != config.ReactionMessageId {
		return
	}

	roleId, exists := config.ReactionRolesMap[r.Emoji.Name]

	if !exists {
		slog.Warn("Role doesn't exist", "name", r.Emoji.Name)
		return
	}

	err := s.GuildMemberRoleAdd(r.GuildID, r.UserID, roleId)

	if err != nil {
		slog.Error("Error adding role", "error", err)
	}

}

func OnMessageReactionRemove(s *discordgo.Session, r *discordgo.MessageReactionRemove) {
	slog.Info("%v remove reaction %v", r.UserID, r.Emoji.Name)

	if r.MessageID != config.ReactionMessageId {
		return
	}

	roleId, exists := config.ReactionRolesMap[r.Emoji.Name]

	if !exists {
		slog.Warn("Role doesn't exist", "name", r.Emoji.Name)
		return
	}

	err := s.GuildMemberRoleRemove(r.GuildID, r.UserID, roleId)

	if err != nil {
		slog.Error("Error removing role", "error", err)
	}
}
