package dto

type UpdateBotSettingsDTO struct {
	BotStatus     string   `json:"bot_status" binding:"required"`
	Description   string   `json:"description"`
	HelpMessage   string   `json:"help_message"`
	HelloMessages []string `json:"hello_messages"`
}

type UpdateGuildSettingsDTO struct {
	ID    string           `json:"id" binding:"required"`
	Roles RolesSettingsDTO `json:"roles"`
}

type RolesSettingsDTO struct {
	MessageID  string            `json:"message_id"`
	Matching   map[string]string `json:"matching"`
	IsDisabled bool              `json:"is_disabled"`
}
