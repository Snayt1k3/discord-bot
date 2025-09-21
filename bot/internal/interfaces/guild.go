package interfaces

import dtoGuild "bot/internal/dto/guild"

type SettingsAdapterInterface interface {
	Get(guildID string) (dtoGuild.GuildSettings, error)
	Create(guildID string) error
}

type WelcomeAdapterInterface interface {
	SetChannel(guildID, channelID string) (error)
	AddMessage(guildID, message string) (error)
	DeleteMessage(guildID, message string) (error)
}

type RolesAdapterInterface interface {
	Add(roleId, emoji, guildID string) (error)
	Delete(roleId, emoji, guildID string) (error)
	SetMessageID(messageID, guildID string) (error)
}

type AutoModeAdapterInterface interface {
	Toggle(guildID string, enabled bool) (error)
	AddBannedWord(guildID, word string) (error)
	RemoveBannedWord(guildID, word string) (error)
	AddCapsLockChannel(guildID, channelID string) (error)
	RemoveCapsLockChannel(guildID, channelID string) (error)
	AddAntiLinkChannel(guildID, channelID string) (error)
	RemoveAntiLinkChannel(guildID, channelID string) (error)
}

type LogAdapterInterface interface {
	Toggle(guildID string, enabled bool) (error)
	AddChannel(guildID, channelID string) (error)
	RemoveChannel(guildID, channelID string) (error)
}