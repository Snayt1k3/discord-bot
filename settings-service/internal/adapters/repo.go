package adapters

import (
	"errors"
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


// Реализация методов для GuildSettingRepository

func (r *GuildSettingRepository) Create(setting *models.GuildSetting) (*models.GuildSetting, error) {
	if err := r.db.Create(setting).Error; err != nil {
		return nil, err
	}
	return setting, nil
}

func (r *GuildSettingRepository) GetByID(id uint) (*models.GuildSetting, error) {
	var setting models.GuildSetting
	if err := r.db.First(&setting, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("guild setting not found")
		}
		return nil, err
	}
	return &setting, nil
}

func (r *GuildSettingRepository) Update(setting *models.GuildSetting) error {
	if err := r.db.Save(setting).Error; err != nil {
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

	// Строим запрос с фильтрацией
	query := r.db.Model(&models.GuildSetting{})

	// Проходим по фильтрам
	for key, value := range filters {
		query = query.Where(key+" = ?", value)  // Фильтруем по ключу и значению
	}

	// Выполняем запрос и получаем результаты
	if err := query.Find(&guildSettings).Error; err != nil {
		return nil, err
	}

	// Возвращаем найденные записи
	return guildSettings, nil
}
// Реализация методов для BotSettingRepository

func (r *BotSettingRepository) Create(setting *models.BotSetting) (*models.BotSetting, error) {
	if err := r.db.Create(setting).Error; err != nil {
		return nil, err
	}
	return setting, nil
}

func (r *BotSettingRepository) Filter(filters map[string]interface{}) ([]models.BotSetting, error) {
	var guildSettings []models.BotSetting

	// Строим запрос с фильтрацией
	query := r.db.Model(&models.BotSetting{})

	// Проходим по фильтрам
	for key, value := range filters {
		query = query.Where(key+" = ?", value)  // Фильтруем по ключу и значению
	}

	// Выполняем запрос и получаем результаты
	if err := query.Find(&guildSettings).Error; err != nil {
		return nil, err
	}

	// Возвращаем найденные записи
	return guildSettings, nil
}

func (r *BotSettingRepository) GetByID(id uint) (*models.BotSetting, error) {
	var setting models.BotSetting
	if err := r.db.First(&setting, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("bot setting not found")
		}
		return nil, err
	}
	return &setting, nil
}

func (r *BotSettingRepository) Update(setting *models.BotSetting) error {
	if err := r.db.Save(setting).Error; err != nil {
		return err
	}
	return nil
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