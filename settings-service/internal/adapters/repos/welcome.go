package repos

import (
	"gorm.io/gorm"
	"settings-service/internal/models"
)


type WelcomeRepositoryImpl struct {
	db *gorm.DB
}

func NewWelcomeRepository(db *gorm.DB) *WelcomeRepositoryImpl {
	return &WelcomeRepositoryImpl{
		db: db,
	}
}

func (r *WelcomeRepositoryImpl) SetWelcomeChannel(guildId, channelId string) error {
	return r.db.Model(&models.WelcomeSettings{}).
		Where("guild_id = ?", guildId).
		Update("channel_id", channelId).Error
}

func (r *WelcomeRepositoryImpl) AddWelcomeMessage(guildId, message string) error {
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

func (r *WelcomeRepositoryImpl) DeleteWelcomeMessage(guildId, message string) error {
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
