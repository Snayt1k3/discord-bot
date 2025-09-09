package interfaces

import dtoGuild "bot/internal/dto/guild"

type GuildServiceInterface interface {
	GetGuildSettings(guildID string) (dtoGuild.GuildSettings, error)
	CreateSettings(guildID string) error

	AddRole(roleId, emoji, guildID string) (dtoGuild.Role, error)
	DeleteRole(roleId, emoji, guildID string) (dtoGuild.Role, error)
	SetRoleMessageID(messageID, guildID string) (dtoGuild.RoleMessage, error)

	SetWelcomeChannel(guildID, channelID string) (dtoGuild.SetWelcomeChannelResponse, error)
	AddWelcomeMessage(guildID, message string) (dtoGuild.WelcomeMessageResponse, error)
	DeleteWelcomeMessage(guildID, message string) (dtoGuild.WelcomeMessageResponse, error)
}
