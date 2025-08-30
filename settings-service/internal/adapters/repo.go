package adapters

import (
	"gorm.io/gorm"
	"settings-service/internal/interfaces"
	"settings-service/internal/models"
)

type GuildRepositoryImpl struct {
	db *gorm.DB
}

func NewGuildRepository(db *gorm.DB) interfaces.GuildRepository {
	return &GuildRepositoryImpl{db: db}
}

func (r *GuildRepositoryImpl) CreateGuildSetting(guildId string) error {
	obj := &models.Settings{
		GuildID: guildId,
		Role: models.RolesSettings{
			GuildID: guildId,
		},
		Welcome: models.WelcomeSettings{
			GuildID: guildId,
		},
	}
	return r.db.Create(obj).Error
}

func (r *GuildRepositoryImpl) GetGuildSettings(guildID string) (*models.Settings, error) {
	var settings models.Settings
	if err := r.db.Preload("Role.Roles").Preload("Welcome.Messages").Where("guild_id = ?", guildID).First(&settings).Error; err != nil {
		return nil, err
	}
	return &settings, nil
}

func (r *GuildRepositoryImpl) DeleteGuildSettings(guildID string) error {
	return r.db.Where("guild_id = ?", guildID).Delete(&models.Settings{}).Error
}

func (r *GuildRepositoryImpl) SetRoleMessageId(guildId, messageId string) error {
	return r.db.Model(&models.RolesSettings{}).
		Where("guild_id = ?", guildId).
		Update("message_id", messageId).Error
}

func (r *GuildRepositoryImpl) AddRole(guildId, roleId, emoji string) error {
	var roleSetting models.RolesSettings
	if err := r.db.Where("guild_id = ?", guildId).First(&roleSetting).Error; err != nil {
		return err
	}

	role := models.Role{
		RoleID: roleId,
		Emoji:  emoji,
	}

	if err := r.db.FirstOrCreate(&role, models.Role{RoleID: roleId, Emoji: emoji}).Error; err != nil {
		return err
	}

	return r.db.Model(&roleSetting).Association("Roles").Append(&role)
}

func (r *GuildRepositoryImpl) DeleteRole(guildId, roleId string) error {
	var roleSetting models.RolesSettings
	if err := r.db.Where("guild_id = ?", guildId).First(&roleSetting).Error; err != nil {
		return err
	}

	var role models.Role
	if err := r.db.Where("role_id = ?", roleId).First(&role).Error; err != nil {
		return err
	}

	if err := r.db.Model(&roleSetting).Association("Roles").Delete(&role); err != nil {
		return err
	}

	var count int64
	r.db.Model(&models.RolesSettings{}).Where("id IN (SELECT roles_settings_id FROM roles_settings_roles WHERE role_id = ?)", role.ID).Count(&count)
	if count == 0 {
		return r.db.Delete(&role).Error
	}

	return nil
}

func (r *GuildRepositoryImpl) SetWelcomeChannel(guildId, channelId string) error {
	return r.db.Model(&models.WelcomeSettings{}).
		Where("guild_id = ?", guildId).
		Update("channel_id", channelId).Error
}

func (r *GuildRepositoryImpl) AddWelcomeMessage(guildId, message string) error {
    var welcomeSetting models.WelcomeSettings
    if err := r.db.Where("guild_id = ?", guildId).First(&welcomeSetting).Error; err != nil {
        return err
    }

    msg := models.Message{Message: message}
    if err := r.db.FirstOrCreate(&msg, models.Message{Message: message}).Error; err != nil {
        return err
    }

    return r.db.Model(&welcomeSetting).Association("Messages").Append(&msg)
}

func (r *GuildRepositoryImpl) DeleteWelcomeMessage(guildId, message string) error {
	var welcomeSetting models.WelcomeSettings
	if err := r.db.Where("guild_id = ?", guildId).First(&welcomeSetting).Error; err != nil {
		return err
	}

	var msg models.Message
	if err := r.db.Where("message = ?", message).First(&msg).Error; err != nil {
		return err
	}

	if err := r.db.Model(&welcomeSetting).Association("Messages").Delete(&msg); err != nil {
		return err
	}

	var count int64
	r.db.Model(&models.WelcomeSettings{}).Where("id IN (SELECT welcome_setting_id FROM welcome_setting_messages WHERE message_id = ?)", msg.ID).Count(&count)
	if count == 0 {
		return r.db.Delete(&msg).Error
	}

	return nil
}
