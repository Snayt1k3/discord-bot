package models


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
