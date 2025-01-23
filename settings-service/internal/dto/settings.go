package dto


type RolesSettings struct {
	MesssageId string `json:"message_id"`
	Matching map[string]string `json:"matching"`
}

type SettingsJson struct {
	// This class is part of GuildSetting Model
	Roles RolesSettings `json:"roles"`
}

func (s *SettingsJson) AddRole(key string, value string) {
	s.Roles.Matching[key] = value
}

func (s *SettingsJson) RemoveRole(key string) {
	delete(s.Roles.Matching, key)
}

func (s *SettingsJson) SetMessageId(messageId string) {
	s.Roles.MesssageId = messageId
}



type BotSettingsDTO struct {
	ID uint `json:id`
	BotStatus string `json:bot_status`
	Description string `json:description`
	HelpMessage string `json:help_message`
	HelloMessages []string `json:hello_messages`
}


type BotSettingsUpdate struct {
	BotStatus string `json:bot_status`
	Description string `json:description`
	HelpMessage string `json:help_message`
	HelloMessages []string `json:hello_messages`
}

type GuildSettingsDTO struct {
	ID uint `json:id`
	GuildID string `json:guild_id`
	Settings SettingsJson `json:settings`
}

type GuildSettingsUpdateDTO struct {
	Roles RolesSettings `json:roles`
}

type GuildSettingsDeleteDTO struct {
	Roles bool `json:roles`
}

