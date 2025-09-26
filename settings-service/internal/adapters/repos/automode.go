package repos

import (
	"errors"

	"gorm.io/gorm"

	"settings-service/internal/models"
)

type AutoModeRepositoryImpl struct {
	db *gorm.DB
}

func NewAutoModeRepository(db *gorm.DB) *AutoModeRepositoryImpl {
	return &AutoModeRepositoryImpl{
		db: db,
	}
}

// Вспомогательный метод – находит или создаёт AutoModeSettings для гильдии
func (r *AutoModeRepositoryImpl) getOrCreateAutoMode(guildId string) (*models.AutoModeSettings, error) {
	var settings models.AutoModeSettings
	err := r.db.Where("guild_id = ?", guildId).First(&settings).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		settings = models.AutoModeSettings{GuildID: guildId, Enabled: false}
		if err := r.db.Create(&settings).Error; err != nil {
			return nil, err
		}
	} else if err != nil {
		return nil, err
	}

	return &settings, nil
}

func (r *AutoModeRepositoryImpl) ToggleAutoMode(guildId string, enabled bool) (models.AutoModeSettings, error) {
	var settings models.AutoModeSettings

	if err := r.db.Model(&settings).
		Where("guild_id = ?", guildId).
		Update("enabled", enabled).Error; err != nil {
		return models.AutoModeSettings{}, err
	}

	if err := r.db.Where("guild_id = ?", guildId).First(&settings).Error; err != nil {
		return models.AutoModeSettings{}, err
	}

	return settings, nil
}

func (r *AutoModeRepositoryImpl) AddBannedWord(guildId string, word string) (models.BannedWord, error) {
	settings, err := r.getOrCreateAutoMode(guildId)
	if err != nil {
		return models.BannedWord{}, err
	}

	bannedWord := models.BannedWord{Word: word}
	if err := r.db.Model(settings).Association("BannedWords").Append(&bannedWord); err != nil {
		return models.BannedWord{}, err
	}
	return bannedWord, nil
}

func (r *AutoModeRepositoryImpl) DeleteBannedWord(guildId string, word string) error {
	settings, err := r.getOrCreateAutoMode(guildId)
	if err != nil {
		return err
	}

	var bw models.BannedWord
	if err := r.db.Where("word = ?", word).First(&bw).Error; err != nil {
		return err
	}
	return r.db.Model(settings).Association("BannedWords").Delete(&bw)
}

func (r *AutoModeRepositoryImpl) AddCapsLockChannel(guildId string, channelId string) (models.AntiCapsChannel, error) {
	settings, err := r.getOrCreateAutoMode(guildId)
	if err != nil {
		return models.AntiCapsChannel{}, err
	}

	model := models.AntiCapsChannel{ChannelID: channelId}
	if err := r.db.Model(settings).Association("CapsLocks").Append(&model); err != nil {
		return models.AntiCapsChannel{}, err
	}
	return model, nil
}


func (r *AutoModeRepositoryImpl) DeleteCapsLockChannel(guildId string, channelId string) error {
	settings, err := r.getOrCreateAutoMode(guildId)
	if err != nil {
		return err
	}

	var ch models.AntiCapsChannel
	if err := r.db.Where("channel_id = ?", channelId).First(&ch).Error; err != nil {
		return err
	}
	return r.db.Model(settings).Association("CapsLocks").Delete(&ch)
}

func (r *AutoModeRepositoryImpl) AddAntiLinkChannel(guildId string, channelId string) (models.AntiLinkChannel, error) {
	settings, err := r.getOrCreateAutoMode(guildId)
	if err != nil {
		return models.AntiLinkChannel{}, err
	}

	model := models.AntiLinkChannel{ChannelID: channelId}
	if err := r.db.Model(settings).Association("AntiLinks").Append(&model); err != nil {
		return models.AntiLinkChannel{}, err
	}
	return model, nil
}

func (r *AutoModeRepositoryImpl) DeleteAntiLinkChannel(guildId string, channelId string) error {
	settings, err := r.getOrCreateAutoMode(guildId)
	if err != nil {
		return err
	}

	var ch models.AntiLinkChannel
	if err := r.db.Where("channel_id = ?", channelId).First(&ch).Error; err != nil {
		return err
	}
	return r.db.Model(settings).Association("AntiLinks").Delete(&ch)
}
