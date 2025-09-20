package repos

import (
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
	bannedWord := models.BannedWord{
		Word:    word,
	}
	if err := r.db.Create(&bannedWord).Error; err != nil {
		return models.BannedWord{}, err
	}
	return bannedWord, nil
}

func (r *AutoModeRepositoryImpl) DeleteBannedWord(guildId string, word string) error {
	res := r.db.Where("word = ?", word).Where("guild_id = ?", guildId).Delete(&models.BannedWord{})
	return res.Error
}

func (r *AutoModeRepositoryImpl) AddCapsLockChannel(guildId string, channelId string) (models.AntiCapsChannel, error) {
	model := models.AntiCapsChannel{
		ChannelID: channelId,
	}
	if err := r.db.Create(&model).Error; err != nil {
		return models.AntiCapsChannel{}, err
	}
	return model, nil
}

func (r *AutoModeRepositoryImpl) DeleteCapsLockChannel(guildId string, channelId string) error {
	return r.db.Where("channel_id = ?", channelId).Where("guild_id = ?", guildId).Delete(models.AntiCapsChannel{}).Error
}

func (r *AutoModeRepositoryImpl) AddAntiLinkChannel(guildId string, channelId string) (models.AntiLinkChannel, error) {
	model := models.AntiLinkChannel{
		ChannelID: channelId,
	}
	if err := r.db.Create(&model).Error; err != nil {
		return models.AntiLinkChannel{}, err
	}
	return model, nil
}

func (r *AutoModeRepositoryImpl) DeleteAntiLinkChannel(guildId string, channelId string) error {
	return r.db.Where("channel_id = ?", channelId).Where("guild_id = ?", guildId).Delete(models.AntiLinkChannel{}).Error
}
