package adapters

import (
	"settings-service/internal/dto"
	"settings-service/internal/interfaces"
	"settings-service/internal/models"
)



type SettingsService struct {
	GuildRepo interfaces.Repository[models.GuildSetting]
	BotRepo   interfaces.Repository[models.BotSetting]
}

// Обновление настроек бота
func (s *SettingsService) UpdateBotSettings(data dto.BotSettingsUpdate) (dto.BotSettingsDTO, error) {
	// Извлекаем текущие настройки бота
	botSetting, err := s.BotRepo.GetByID(1) // Предположим, что для бота используется ID = 1
	if err != nil {
		return dto.BotSettingsDTO{}, err
	}

	// Обновляем поля бота
	botSetting.BotStatus = data.BotStatus
	botSetting.Description = data.Description
	botSetting.HelpMessage = data.HelpMessage
	botSetting.HelloMessages = data.HelloMessages

	// Сохраняем обновленные настройки
	err = s.BotRepo.Update(botSetting)
	if err != nil {
		return dto.BotSettingsDTO{}, err
	}

	// Возвращаем DTO с обновленными данными
	return dto.BotSettingsDTO{
		ID:           botSetting.ID,
		BotStatus:    botSetting.BotStatus,
		Description:  botSetting.Description,
		HelpMessage:  botSetting.HelpMessage,
		HelloMessages: botSetting.HelloMessages,
	}, nil
}

// Получение настроек бота
func (s *SettingsService) GetBotSettings() (dto.BotSettingsDTO, error) {
	// Извлекаем настройки бота по ID
	botSetting, err := s.BotRepo.GetByID(1) // Используем ID = 1 для примера
	if err != nil {
		return dto.BotSettingsDTO{}, err
	}

	// Возвращаем DTO с настройками
	return dto.BotSettingsDTO{
		ID:           botSetting.ID,
		BotStatus:    botSetting.BotStatus,
		Description:  botSetting.Description,
		HelpMessage:  botSetting.HelpMessage,
		HelloMessages: botSetting.HelloMessages,
	}, nil
}

// Получение настроек гильдии по ID
func (s *SettingsService) GetByGuildID(id int) (dto.GuildSettingsDTO, error) {
	// Извлекаем настройки гильдии по ID
	guildSetting, err := s.GuildRepo.GetByID(uint(id))
	if err != nil {
		return dto.GuildSettingsDTO{}, err
	}

	// Возвращаем DTO с настройками для гильдии
	return dto.GuildSettingsDTO{
		ID:        guildSetting.ID,
		GuildID:   guildSetting.GuildID,
		Settings:  guildSetting.Settings,
	}, nil
}

// Получение всех настроек гильдий
func (s *SettingsService) GetAllGuildSettings() ([]dto.GuildSettingsDTO, error) {
	// Извлекаем все настройки гильдий
	guildSettings, err := s.GuildRepo.GetAll()
	if err != nil {
		return nil, err
	}

	// Преобразуем данные в DTO
	var result []dto.GuildSettingsDTO
	for _, setting := range guildSettings {
		result = append(result, dto.GuildSettingsDTO{
			ID:        setting.ID,
			GuildID:   setting.GuildID,
			Settings:  setting.Settings,
		})
	}

	return result, nil
}

// Обновление настроек гильдии
func (s *SettingsService) UpdateGuildSettings(id int, data dto.GuildSettingsUpdateDTO) (dto.GuildSettingsDTO, error) {
	// Извлекаем настройки гильдии по ID
	guildSetting, err := s.GuildRepo.GetByID(uint(id))
	if err != nil {
		return dto.GuildSettingsDTO{}, err
	}

	// Обновляем поля
	guildSetting.Settings.Roles = data.Roles

	// Сохраняем обновленные настройки
	err = s.GuildRepo.Update(guildSetting)
	if err != nil {
		return dto.GuildSettingsDTO{}, err
	}

	// Возвращаем DTO с обновленными данными
	return dto.GuildSettingsDTO{
		ID:        guildSetting.ID,
		GuildID:   guildSetting.GuildID,
		Settings:  guildSetting.Settings,
	}, nil
}

// Удаление настроек гильдии
func (s *SettingsService) DeleteGuildSetting(id int, data dto.GuildSettingsDeleteDTO) error {
	// Извлекаем настройки гильдии по ID
	guildSetting, err := s.GuildRepo.GetByID(uint(id))
	if err != nil {
		return err
	}

	// Удаляем данные в зависимости от запроса
	if data.Roles {
		guildSetting.Settings.Roles = dto.RolesSettings{}
	}

	// Сохраняем изменения
	err = s.GuildRepo.Update(guildSetting)
	if err != nil {
		return err
	}

	return nil
}