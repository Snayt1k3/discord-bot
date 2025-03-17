package interfaces

import (
	"settings-service/internal/dto"
)

type SettingsService interface {
	GetSettingsByGuildID(guildID string) (*dto.GuildSettingsDTO, error)
	CreateGuildSettings(guildID string) error
	UpdateRolesSettings(roles *dto.RolesSettings) (*dto.GuildSettingsDTO, error)
}
