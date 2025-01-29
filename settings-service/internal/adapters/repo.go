package adapters

import (
	"settings-service/internal/models"
	"gorm.io/gorm"
)


type GuildSettingRepository struct {
	db *gorm.DB
}

type BotSettingRepository struct {
	db *gorm.DB
}

// Конструкторы для репозиториев
func NewGuildSettingRepository(db *gorm.DB) *GuildSettingRepository {
	return &GuildSettingRepository{db: db}
}

func NewBotSettingRepository(db *gorm.DB) *BotSettingRepository {
	return &BotSettingRepository{db: db}
}



func (r *GuildSettingRepository) Create(setting *models.GuildSetting) (*models.GuildSetting, error) {
	if err := r.db.Create(setting).Error; err != nil {
		return nil, err
	}
	return setting, nil
}

func (r *GuildSettingRepository) Updates(id string, fields map[string]interface{}) error {
	for key, value := range fields {
		if value == nil || value == "" {
			delete(fields, key)
		}
	}

	var setting models.GuildSetting
	if err := r.db.First(&setting, "id = ?", id).Error; err != nil {
		return err
	}

	if err := r.db.Model(&setting).Updates(fields).Error; err != nil {
		return err 
	}

	return nil
}

func (r *GuildSettingRepository) Delete(id uint) error {
	if err := r.db.Delete(&models.GuildSetting{}, id).Error; err != nil {
		return err
	}
	return nil
}

func (r *GuildSettingRepository) GetAll() ([]models.GuildSetting, error) {
	var settings []models.GuildSetting
	if err := r.db.Find(&settings).Error; err != nil {
		return nil, err
	}
	return settings, nil
}

func (r *GuildSettingRepository) Filter(filters map[string]interface{}) ([]models.GuildSetting, error) {
	var guildSettings []models.GuildSetting

	query := r.db.Model(&models.GuildSetting{})

	for key, value := range filters {
		query = query.Where(key+" = ?", value)
	}

	if err := query.Find(&guildSettings).Error; err != nil {
		return nil, err
	}

	return guildSettings, nil
}


func (r *BotSettingRepository) Create(setting *models.BotSetting) (*models.BotSetting, error) {
	if err := r.db.Create(setting).Error; err != nil {
		return nil, err
	}
	return setting, nil
}

func (r *BotSettingRepository) Updates(id string, fields map[string]interface{}) error {
	for key, value := range fields {
		if value == nil || value == "" {
			delete(fields, key)
		}
	}


	var setting models.BotSetting
	if err := r.db.First(&setting, "id = ?", 1).Error; err != nil {
		return err
	}

	if err := r.db.Model(&setting).Updates(fields).Error; err != nil {
		return err 
	}
	
	return nil
}

func (r *BotSettingRepository) Filter(filters map[string]interface{}) ([]models.BotSetting, error) {
	var guildSettings []models.BotSetting
	query := r.db.Model(&models.BotSetting{})

	for key, value := range filters {
		query = query.Where(key+" = ?", value)  
	}

	if err := query.Find(&guildSettings).Error; err != nil {
		return nil, err
	}

	return guildSettings, nil
}


func (r *BotSettingRepository) Delete(id uint) error {
	if err := r.db.Delete(&models.BotSetting{}, id).Error; err != nil {
		return err
	}
	return nil
}

func (r *BotSettingRepository) GetAll() ([]models.BotSetting, error) {
	var settings []models.BotSetting
	if err := r.db.Find(&settings).Error; err != nil {
		return nil, err
	}
	return settings, nil
}