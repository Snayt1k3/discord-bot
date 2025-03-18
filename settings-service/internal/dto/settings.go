package dto

type GuildSettingsDTO struct {
	ID      uint          `json:"id"`
	GuildID string        `json:"guild_id"`
	Roles   RolesSettings `json:"roles"`
	Welcome WelcomeSettings `json:"welcome"`
}
