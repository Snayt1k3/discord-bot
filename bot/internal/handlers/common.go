package handlers

import (
	"bot/internal/dto/discord"
	"bot/internal/handlers/gachas"
	"bot/internal/handlers/settings"
)

var (
	userHandlers = map[string]func(data discord.HandlerData) error{
		"play":   PlayCommandHandler,
		"skip":   SkipCommandHandler,
		"stop":   StopCommandHandler,
		"help":   HelpHandler,
		"resume": ResumeHandler,
		"gachas": gachas.ShowGachas,
	}
	adminHandlers = map[string]func(data discord.HandlerData){
		"add-role-reactions":    settings.AddRole,
		"remove-role-reactions": settings.RemoveRole,
		"set-message-id":        settings.SetMessageId,
		"set-welcome-channel":   settings.SetChannelId,
		"settings":              SettingsHandler,
	}
	buttons = map[string]func(data discord.HandlerData){
		"view_reaction_roles": settings.ShowAllRoles,
	}
)
