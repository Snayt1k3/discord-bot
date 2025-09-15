package interfaces

import "settings-service/internal/models"

type AutoModeRepository interface {
	// Переключение автомода (вкл/выкл)
	ToggleAutoMode(guildId string, enabled bool) (models.AutoModeSettings, error)

	// Работа со стоп-словами
	AddBannedWord(guildId string, word string) (models.BannedWord, error)
	DeleteBannedWord(guildId string, word string) error

	// Работа с антикапсом
	AddCapsLockChannel(guildId string, channelId string) (models.AntiCapsChannel, error)
	DeleteCapsLockChannel(guildId string, channelId string) error

	// Работа с антилинками
	AddAntiLinkChannel(guildId string, channelId string) (models.AntiLinkChannel, error)
	DeleteAntiLinkChannel(guildId string, channelId string) error
}

type LogRepository interface {
	// Добавить канал логов
	AddLogChannel(guildId string, channelId string) (models.LogSettings, error)

	// Удалить канал логов
	RemoveLogChannel(guildId string, channelId string) error

	// Включить/выключить логи
	ToggleLog(guildId string, enabled bool) error
}

type ReactionRolesRepository interface {
	// Задать messageId для реакционных ролей
	SetRoleMessageId(guildId, messageId string) error

	// Добавить роль по emoji
	AddRole(guildId, roleId, emoji string) (models.Role, error)

	// Удалить роль
	DeleteRole(guildId, roleId string) error
}

type GuildSettingsRepository interface {
	// Создать настройки гильдии
	CreateGuildSetting(guildId string) (models.Settings, error)

	// Получить настройки гильдии с прелоадами
	GetGuildSettings(guildID string) (*models.Settings, error)
}


type WelcomeRepository interface {
	// Задать канал для welcome-сообщений
	SetWelcomeChannel(guildId, channelId string) error

	// Добавить welcome-сообщение
	AddWelcomeMessage(guildId, message string) error

	// Удалить welcome-сообщение
	DeleteWelcomeMessage(guildId, message string) error
}