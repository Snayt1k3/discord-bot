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

func (r *LogRepositoryImpl) AddLogs(guildId, channelId string, events []models.EventType) (models.LogSettings, error) {
	var log models.LogSettings

	if err := r.db.
		Preload("Events").
		Where("guild_id = ?", guildId).
		First(&log).Error; err != nil {
		return models.LogSettings{}, err
	}

	exists := make(map[string]struct{})

	for _, e := range log.Events {
		key := string(e.EventType) + ":" + e.ChannelID
		exists[key] = struct{}{}
	}

	for _, eventType := range events {
		key := string(eventType) + ":" + channelId

		if _, ok := exists[key]; ok {
			continue
		}

		log.Events = append(log.Events, models.LogEvent{
			EventType:     eventType,
			ChannelID:    channelId,
		})
	}

	if err := r.db.Save(&log).Error; err != nil {
		return models.LogSettings{}, err
	}

	return log, nil
}

func (r *LogRepositoryImpl) RemoveLogs(guildId, channelId string, events []models.EventType) error {
	var log models.LogSettings

	if err := r.db.
		Where("guild_id = ?", guildId).
		First(&log).Error; err != nil {
		return err
	}

	if len(events) == 0 {
		return nil
	}

	if err := r.db.
		Where(
			"log_settings_id = ? AND channel_id = ? AND event_type IN ?",
			log.ID,
			channelId,
			events,
		).
		Delete(&models.LogEvent{}).Error; err != nil {
		return err
	}

	return nil
}

func (r *LogRepositoryImpl) ToggleLog(guildId string, enabled bool) error {
	return r.db.Model(&models.LogSettings{}).
		Where("guild_id = ?", guildId).
		Update("enabled", enabled).Error
}
