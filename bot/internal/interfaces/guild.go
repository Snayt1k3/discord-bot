package interfaces

import "bot/internal/dto"

type GuildKeeperInterface interface {
	GetGuildSettings(guildId string) (dto.GuildSettingsResponse, error)
	CreateSettings(guild_id string) error
	UpdateGuildSettings(guildId string, roles dto.RolesSettings) (dto.GuildSettingsResponse, error)
}
