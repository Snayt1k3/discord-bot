package repos

import (
	"settings-service/internal/models"

	"gorm.io/gorm"
)

type LogRepositoryImpl struct {
	db *gorm.DB
}

func NewLogRepository(db *gorm.DB) *LogRepositoryImpl {
	return &LogRepositoryImpl{
		db: db,
	}
}

func (r *LogRepositoryImpl) AddLogChannel(guildId string, channelId string) (models.LogSettings, error) {
	model := &models.LogSettings{
		GuildID:   guildId,
		ChannelID: channelId,
	}

	if err := r.db.Create(model).Error; err != nil {
		return models.LogSettings{}, err
	}
	return *model, nil
}

func (r *LogRepositoryImpl) RemoveLogChannel(guildId string, channelId string) error {
	return r.db.Where("channel_id = ?", channelId).Where("guild_id = ?", guildId).Delete(models.LogSettings{}).Error
}

func (r *LogRepositoryImpl) ToggleLog(guildId string, enabled bool) error {
	return r.db.Model(&models.LogSettings{}).
		Where("guild_id = ?", guildId).
		Update("enabled", enabled).Error
}
