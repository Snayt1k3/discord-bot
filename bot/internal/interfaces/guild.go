package interfaces

import dtoGuild "bot/internal/dto/settings"

type GuildKeeperInterface interface {
	GetGuildSettings(guildId string) (dtoGuild.GuildSettingsResponse, error)
	CreateSettings(guild_id string) error
	UpdateRolesSetting(guildId string, roles dtoGuild.RolesSettings) (dtoGuild.GuildSettingsResponse, error)
	UpdateWelcomeSetting(guildId string, welcome dtoGuild.WelcomeSettings) error
}
