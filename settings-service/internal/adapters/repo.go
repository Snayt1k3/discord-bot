package adapters

import (
	"encoding/json"
	"fmt"
	"settings-service/internal/dto"
	"settings-service/internal/interfaces"
	"settings-service/internal/models"

	"gorm.io/gorm"
)

type GuildRepositoryImpl struct {
	db *gorm.DB
}

func NewGuildRepository(db *gorm.DB) interfaces.GuildRepository {
	return &GuildRepositoryImpl{db: db}
}

func (r *GuildRepositoryImpl) CreateGuildSetting(guild *models.GuildSetting) error {
	return r.db.Create(guild).Error
}

func (r *GuildRepositoryImpl) GetGuildSetting(guildID string) (*models.GuildSetting, error) {
	var guild models.GuildSetting
	err := r.db.Preload("Role").Preload("Welcome").Where("guild_id = ?", guildID).First(&guild).Error
	if err != nil {
		return nil, err
	}
	return &guild, nil
}

func (r *GuildRepositoryImpl) PatchGuildSetting(guildID string, updates map[string]interface{}) error {
	return r.db.Model(&models.GuildSetting{}).Where("guild_id = ?", guildID).Updates(updates).Error
}

func (r *GuildRepositoryImpl) DeleteGuildSetting(guildID string) error {
	return r.db.Where("guild_id = ?", guildID).Delete(&models.GuildSetting{}).Error
}

func (r *GuildRepositoryImpl) UpdateRoleSetting(role *dto.RolesSettings) error {
	updates := map[string]interface{}{}

	if role.MessageId != "" {
		updates["message_id"] = role.MessageId
	}
	fmt.Println(role.Matching)
	if role.Matching != nil {
		roleBytes, err := json.Marshal(role.Matching)
		if err != nil {
			return err
		}
		updates["role"] = roleBytes
	}

	if len(updates) == 0 {
		return nil
	}

	return r.db.Model(&models.RoleSetting{}).Where("guild_id = ?", role.GuildID).Updates(updates).Error
}

func (r *GuildRepositoryImpl) UpdateWelcomeSetting(welcome *dto.WelcomeSettings) error {
	updates := map[string]any{};{}

	if welcome.ChannelId != "" {
		updates["channel_id"] = welcome.ChannelId
	}

	if len(updates) == 0 {
		return nil
	}

	return r.db.Model(&models.WelcomeSetting{}).Where("guild_id = ?", welcome.GuildID).Updates(updates).Error
}