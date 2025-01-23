package interfaces

import (
	"settings-service/internal/dto"
)

type SettingsService interface {
	// bot settings methods
	UpdateBotSettings(data dto.BotSettingsUpdate) (dto.BotSettingsDTO, error)
	GetBotSettings()(dto.BotSettingsDTO, error)

	// guild settings methods
	GetByGuildID(id int) (dto.GuildSettingsDTO, error)
	GetAllGuildSettings() (dto.GuildSettingsDTO, error)
	UpdateGuildSettings(id int, data dto.GuildSettingsUpdateDTO) (dto.GuildSettingsDTO, error)
	DeteleGuildSetting(id int, data dto.GuildSettingsDeleteDTO)

}