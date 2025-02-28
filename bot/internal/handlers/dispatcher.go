package handlers

import (
	"bot/internal/interfaces"
)

type CommandsDispatcher struct {
	settingsService interfaces.ServiceSettingsInterface
}

// todo: add command filter
func (cd *CommandsDispatcher) Dispatch(name string) {}