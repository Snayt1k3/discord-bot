package dto


type RolesSettings struct {
	MesssageId string `json:"message_id"`
	Matching map[string]string `json:"matching"`
	IsDisabled bool `json:"is_disabled"`
}

type SettingsJson struct {
	Roles RolesSettings `json:"roles"`
}


type BotSettingsDTO struct {
	ID uint `json:id`
	BotStatus string `json:"bot_status"`
	Description string `json:"description"`
	HelpMessage string `json:"help_message"`
	HelloMessages []string `json:"hello_messages"`
}


type BotSettingsUpdate struct {
	BotStatus string `json:"bot_status"`
	Description string `json:"description"`
	HelpMessage string `json:"help_message"`
	HelloMessages []string `json:"hello_messages"`
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


