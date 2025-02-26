package adapters

import (
	"encoding/json"
	"settings-service/internal/dto"
	"settings-service/internal/interfaces"
	"settings-service/internal/models"
)

type SettingsService struct {
	GuildRepo interfaces.Repository[models.GuildSetting]
}

func (s *SettingsService) GetByGuildID(id string) (dto.GuildSettingsDTO, error) {
	guildSetting, err := s.GuildRepo.Filter(map[string]interface{}{"guild_id": id})
	if err != nil || len(guildSetting) == 0 {
		return dto.GuildSettingsDTO{}, err
	}

	return dto.GuildSettingsDTO{
		ID:      guildSetting[0].ID,
		GuildID: guildSetting[0].GuildID,
		Roles: dto.RolesSettings{
			MesssageId: guildSetting[0].Roles.MesssageId,
			Matching:   guildSetting[0].Roles.Matching,
		},
	}, nil
}

func (s *SettingsService) UpdateGuildSettings(id string, data dto.GuildSettingsUpdateDTO) (dto.GuildSettingsDTO, error) {
	// Получаем текущие настройки гильдии
	rolesJSON, err := json.Marshal(data.Roles)
	if err != nil {
		return dto.GuildSettingsDTO{}, err
	}

	err = s.GuildRepo.Updates(id, map[string]interface{}{
		"roles": string(rolesJSON),
	})

	if err != nil {
		return dto.GuildSettingsDTO{}, err
	}

	setting, _ := s.GuildRepo.Filter(map[string]interface{}{"guild_id": id})

	return dto.GuildSettingsDTO{
		ID:      setting[0].ID,
		GuildID: setting[0].GuildID,
		Roles: dto.RolesSettings{
			MesssageId: setting[0].Roles.MesssageId,
			Matching:   setting[0].Roles.Matching,
		},
	}, nil
}

func (s *SettingsService) CreateGuildSetting(data dto.GuildSettingsCreateDTO) error {
	model := &models.GuildSetting{GuildID: data.GuildId}

	_, err := s.GuildRepo.Create(model)

	if err != nil {
		return err
	}

	return nil
}
