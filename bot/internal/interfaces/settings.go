package interfaces

import "bot/internal/dto"

type ServiceSettingsInterface interface {
	GetGuildSettings(guildId string) (dto.GuildSettingsDTO, error)
	CreateSettings(guild_id string) error
	UpdateGuildSettings(guildId string, roles dto.RolesSettings) (dto.GuildSettingsDTO, error)
}
