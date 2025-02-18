package interfaces

import (
	"settings-service/internal/dto"
)

type SettingsInterface interface {
	// guild settings methods
	GetByGuildID(id string) (dto.GuildSettingsDTO, error)
	GetAllGuildSettings() ([]dto.GuildSettingsDTO, error)
	UpdateGuildSettings(id string, data dto.GuildSettingsUpdateDTO) (dto.GuildSettingsDTO, error)
	DeleteGuildSetting(id string) (error)

}