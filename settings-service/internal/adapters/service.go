package adapters

import (
	"settings-service/internal/dto"
	"settings-service/internal/interfaces"
	"settings-service/internal/models"

)

type SettingsService struct {
	GuildRepo interfaces.Repository[models.GuildSetting]

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
		Roles: dto.RolesSettings{
			MesssageId: guildSetting[0].Roles.MesssageId,
			Matching:   guildSetting[0].Roles.Matching,
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
			Roles: dto.RolesSettings{
				MesssageId: guildSetting.Roles.MesssageId,
				Matching:   guildSetting.Roles.Matching,
			},

		})
	}
	return result, nil
}

// Обновление настроек гильдии
func (s *SettingsService) UpdateGuildSettings(id string, data dto.GuildSettingsUpdateDTO) (dto.GuildSettingsDTO, error) {
	// Получаем текущие настройки гильдии
	err := s.GuildRepo.Updates(id, map[string]interface{}{
		"roles": data.Roles,
	})

	if err != nil {
		return dto.GuildSettingsDTO{}, err
	}

	setting, _ := s.GuildRepo.Filter(map[string]interface{}{"guild_id": id})

	// Возвращаем обновленные данные
	return dto.GuildSettingsDTO{
		ID:      setting[0].ID,
		GuildID: setting[0].GuildID,
		Roles: dto.RolesSettings{
			MesssageId: setting[0].Roles.MesssageId,
			Matching:   setting[0].Roles.Matching,
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