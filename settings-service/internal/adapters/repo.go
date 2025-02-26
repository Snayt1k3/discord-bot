package adapters

import (
	"gorm.io/gorm"
	"settings-service/internal/models"
)

type GuildSettingRepository struct {
	db *gorm.DB
}

// Конструкторы для репозиториев
func NewGuildSettingRepository(db *gorm.DB) *GuildSettingRepository {
	return &GuildSettingRepository{db: db}
}

func (r *GuildSettingRepository) Create(setting *models.GuildSetting) (*models.GuildSetting, error) {
	if err := r.db.Create(setting).Error; err != nil {
		return nil, err
	}
	return setting, nil
}

func (r *GuildSettingRepository) Updates(id string, fields map[string]interface{}) error {
	for key, value := range fields {
		if value == nil {
			delete(fields, key)
			continue
		}

		if str, ok := value.(string); ok && str == "" {
			delete(fields, key)
		}
	}

	var setting models.GuildSetting
	if err := r.db.First(&setting, "guild_id = ?", id).Error; err != nil {
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
