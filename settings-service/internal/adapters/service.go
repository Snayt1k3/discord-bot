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
	// Получаем текущие настройки бота
	botSetting, err := s.BotRepo.GetByID(1) // Предположим, что у нас есть один экземпляр настроек бота
	if err != nil {
		return dto.BotSettingsDTO{}, err
	}

	// Обновляем настройки
	botSetting.BotStatus = data.BotStatus
	botSetting.Description = data.Description
	botSetting.HelpMessage = data.HelpMessage
	botSetting.HelloMessages = data.HelloMessages

	// Сохраняем обновленные настройки
	err = s.BotRepo.Update(botSetting)
	if err != nil {
		return dto.BotSettingsDTO{}, err
	}

	// Возвращаем обновленные данные
	return dto.BotSettingsDTO{
		ID:            botSetting.ID,
		BotStatus:     botSetting.BotStatus,
		Description:   botSetting.Description,
		HelpMessage:   botSetting.HelpMessage,
		HelloMessages: botSetting.HelloMessages,
	}, nil
}

// Получение текущих настроек бота
func (s *SettingsService) GetBotSettings() (dto.BotSettingsDTO, error) {
	// Получаем текущие настройки бота
	botSetting, err := s.BotRepo.GetByID(1) // Предположим, что у нас есть один экземпляр настроек бота
	if err != nil {
		return dto.BotSettingsDTO{}, err
	}

	// Возвращаем данные
	return dto.BotSettingsDTO{
		ID:            botSetting.ID,
		BotStatus:     botSetting.BotStatus,
		Description:   botSetting.Description,
		HelpMessage:   botSetting.HelpMessage,
		HelloMessages: botSetting.HelloMessages,
	}, nil
}

// Получение настроек гильдии по GuildID
func (s *SettingsService) GetByGuildID(id string) (dto.GuildSettingsDTO, error) {
	// Получаем настройки гильдии
	guildSetting, err := s.GuildRepo.Filter(map[string]interface{}{"guild_id": id})
	if err != nil || len(guildSetting) == 0{
		return dto.GuildSettingsDTO{}, err
	}

	// Возвращаем данные
	return dto.GuildSettingsDTO{
		ID:      guildSetting[0].ID,
		GuildID: guildSetting[0].GuildID,
		Settings: dto.SettingsJson{
			Roles: dto.RolesSettings{
				MesssageId: guildSetting[0].Settings.Roles.MesssageId,
				Matching:   guildSetting[0].Settings.Roles.Matching,
			},
		},
	}, nil
}

// Получение всех настроек гильдий
func (s *SettingsService) GetAllGuildSettings() ([]dto.GuildSettingsDTO, error) {
	// Получаем все настройки гильдий
	guildSettingsList, err := s.GuildRepo.GetAll()
	if err != nil {
		return nil, err
	}

	// Преобразуем в DTO и возвращаем
	var result []dto.GuildSettingsDTO
	for _, guildSetting := range guildSettingsList {
		result = append(result, dto.GuildSettingsDTO{
			ID:      guildSetting.ID,
			GuildID: guildSetting.GuildID,
			Settings: dto.SettingsJson{
				Roles: dto.RolesSettings{
					MesssageId: guildSetting.Settings.Roles.MesssageId,
					Matching:   guildSetting.Settings.Roles.Matching,
				},
			},
		})
	}
	return result, nil
}

// Обновление настроек гильдии
func (s *SettingsService) UpdateGuildSettings(id string, data dto.GuildSettingsUpdateDTO) (dto.GuildSettingsDTO, error) {
	// Получаем текущие настройки гильдии
	guildSettings, err := s.GuildRepo.Filter(map[string]interface{}{"guild_id": id})
	
	if err != nil || len(guildSettings) == 0{
		return dto.GuildSettingsDTO{}, err
	}
	
	firstGuildSetting := guildSettings[0]
	// Обновляем настройки
	firstGuildSetting.Settings.Roles.MesssageId = data.Roles.MesssageId
	firstGuildSetting.Settings.Roles.Matching = data.Roles.Matching

	// Сохраняем обновленные настройки
	err = s.GuildRepo.Update(&firstGuildSetting)
	if err != nil {
		return dto.GuildSettingsDTO{}, err
	}

	// Возвращаем обновленные данные
	return dto.GuildSettingsDTO{
		ID:      firstGuildSetting.ID,
		GuildID: firstGuildSetting.GuildID,
		Settings: dto.SettingsJson{
			Roles: dto.RolesSettings{
				MesssageId: firstGuildSetting.Settings.Roles.MesssageId,
				Matching:   firstGuildSetting.Settings.Roles.Matching,
			},
		},
	}, nil
}

// Удаление настроек гильдии
func (s *SettingsService) DeleteGuildSetting(id string) error {
	// Получаем текущие настройки гильдии
	guildSettings, err := s.GuildRepo.Filter(map[string]interface{}{"guild_id": id})
	if err != nil || len(guildSettings) == 0{
		return err
	}

	// Удаляем настройки
	err = s.GuildRepo.Delete(guildSettings[0].ID)
	if err != nil {
		return err
	}

	return nil
}