package interfaces

import (
	"settings-service/internal/models"
)

type GuildRepository interface {
	GetGuildSettings(guildID string) (*models.Settings, error)
	CreateGuildSetting(guildId string) error
	DeleteGuildSettings(guildID string) error

	SetRoleMessageId(guildId, messageId string) error
	AddRole(guildId, roleId, emoji string) error
	DeleteRole(guildId, roleId, emoji string) error

	SetWelcomeChannel(guildId, channelId string) error
	AddWelcomeMessage(guildId, message string) error
	DeleteWelcomeMessage(guildId, message string) error
}
