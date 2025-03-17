package adapters

import (
	"encoding/json"
	"settings-service/internal/dto"
	"settings-service/internal/interfaces"
	"settings-service/internal/models"
)

type SettingsService struct {
	GuildRepo interfaces.GuildRepository
}

func (s *SettingsService) GetSettingsByGuildID(id string) (*dto.GuildSettingsDTO, error) {
	guildSetting, err := s.GuildRepo.GetGuildSetting(id)

	if err != nil {
		return &dto.GuildSettingsDTO{}, err
	}

	var roles map[string]string
	err = json.Unmarshal([]byte(guildSetting.Role.Role), &roles)

	if err != nil {
		return &dto.GuildSettingsDTO{}, err
	}

	return &dto.GuildSettingsDTO{
		ID:      guildSetting.ID,
		GuildID: guildSetting.GuildID,
		Roles: dto.RolesSettings{
			MesssageId: guildSetting.Role.MessageID,
			Matching:   roles,
		},
	}, nil
}

func (s *SettingsService) CreateGuildSettings(guildID string) error {
	model := &models.GuildSetting{
		GuildID: guildID,
		Role: models.RoleSetting{
			MessageID: "",
			Role:      json.RawMessage{},
		},
		Welcome: models.WelcomeSetting{
			MessageID: "",
			Messages:  json.RawMessage{},
		},
	}

	err := s.GuildRepo.CreateGuildSetting(model)

	if err != nil {
		return err
	}

	return nil
}

func (s *SettingsService) UpdateRolesSettings(roles *dto.RolesSettings) (*dto.GuildSettingsDTO, error) {
	err := s.GuildRepo.UpdateRoleSetting(roles)

	if err != nil {
		return &dto.GuildSettingsDTO{}, err
	}

	return s.GetSettingsByGuildID(roles.GuildID)
}
