package guild

import (
	"bot/internal/interfaces"
)

type GuildAdapter struct {
	Welcome        interfaces.WelcomeAdapterInterface
	RolesReactions interfaces.RolesAdapterInterface
	AutoMode       interfaces.AutoModeAdapterInterface
	Settings       interfaces.SettingsAdapterInterface
	Log            interfaces.LogAdapterInterface
	Interaction    interfaces.InteractionAdapterInterface
}

func NewGuildAdapter(http interfaces.HttpClient) *GuildAdapter {
	return &GuildAdapter{
		Welcome:        NewWelcomeAdapter(http),
		RolesReactions: NewRolesAdapter(http),
		AutoMode:       NewAutoMode(http),
		Settings:       NewSettingsAdapter(http),
		Log:            NewLogAdapter(http),
		Interaction:    NewInteraction(http),
	}
}
