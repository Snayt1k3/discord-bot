package dto


type RolesSettings struct {
	MesssageId string `json:"message_id"`
	Matching map[string]string `json:"matching"`
}

type SettingsJson struct {
	Roles RolesSettings `json:"roles"`
}

type GuildSettingsDTO struct {
	ID uint `json:id`
	GuildID string `json:"guild_id"`
	Roles RolesSettings `json:"roles"`
}

type GuildSettingsUpdateDTO struct {
	GuildId string `json:"guild_id"`
	Roles RolesSettings `json:"roles"`
}

type GuildSettingsCreateDTO struct {
	GuildId string `json:"guild_id"`
}

