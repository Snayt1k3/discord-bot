package repos

import (
	"gachas-service/internal/adapters/storage/models/wuwa"
	"gachas-service/internal/interfaces"
	"gorm.io/gorm"
)

type WuwaRepository struct {
	db *gorm.DB
}

// Получение списка персонажей (короткая версия, но вернём всю сущность — снаружи можно преобразовать в short DTO)
func (r *WuwaRepository) GetCharacters() ([]wuwa.Character, error) {
	var characters []wuwa.Character
	err := r.db.Find(&characters).Error
	return characters, err
}

// Получение персонажа по ID с зависимостями
func (r *WuwaRepository) GetCharacterByID(id string) (wuwa.Character, error) {
	var character wuwa.Character

	err := r.db.
		Preload("Weapon").
		Preload("Ascension.MobMaterial").
		Preload("Talents.DungeonMaterial").
		Preload("Talents.MobMaterial").
		First(&character, "id = ?", id).Error

	return character, err
}

// Получение билда по ID персонажа
func (r *WuwaRepository) GetCharacterBuild(id string) (wuwa.Build, error) {
	var build wuwa.Build
	err := r.db.
		Preload("Character").
		Preload("Weapons").
		Preload("Echoes").
		Preload("Stats").
		Where("character_id = ?", id).
		First(&build).Error

	return build, err
}

// Создание нового репозитория
func NewWuwaRepository(db *gorm.DB) interfaces.RepoInterface[wuwa.Character, wuwa.Build] {
	return &WuwaRepository{
		db: db,
	}
}