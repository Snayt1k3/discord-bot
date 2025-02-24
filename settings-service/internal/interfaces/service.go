package interfaces

import (
	"settings-service/internal/dto"
)

type SettingsInterface interface {
	// guild settings methods
	CreateGuildSetting(data dto.GuildSettingsCreateDTO) error
	GetByGuildID(id string) (dto.GuildSettingsDTO, error)
	UpdateGuildSettings(id string, data dto.GuildSettingsUpdateDTO) (dto.GuildSettingsDTO, error)
}