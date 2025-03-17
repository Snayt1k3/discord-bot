package interfaces

import (
	"settings-service/internal/dto"
	"settings-service/internal/models"
)

type GuildRepository interface {
	CreateGuildSetting(guild *models.GuildSetting) error
	GetGuildSetting(guildID string) (*models.GuildSetting, error)
	PatchGuildSetting(guildID string, updates map[string]interface{}) error
	DeleteGuildSetting(guildID string) error
	UpdateRoleSetting(role *dto.RolesSettings) error
}
