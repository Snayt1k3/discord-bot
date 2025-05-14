package repos

import (
	"gachas-service/internal/adapters/storage/models/wuwa"
	"gachas-service/internal/interfaces"

	"gorm.io/gorm"
)

type WuwaRepository struct {
	db *gorm.DB
}

func (r *WuwaRepository) GetCharacters() ([]wuwa.WuwaCharacter, error) {
	var characters []wuwa.WuwaCharacter
	err := r.db.Find(&characters).Error
	return characters, err
}


func (r *WuwaRepository) GetCharacterByID(id string) (wuwa.WuwaCharacter, error) {
	var character wuwa.WuwaCharacter

	err := r.db.
		Preload("Ascension.MobMaterial").
		Preload("Talents.DungeonMaterial").
		Preload("Talents.MobMaterial").
		First(&character, "id = ?", id).Error

	return character, err
}


func (r *WuwaRepository) GetCharacterBuild(id string) (wuwa.WuwaBuild, error) {
	var build wuwa.WuwaBuild
	err := r.db.
		Preload("Character").
		Preload("Character.Ascension.MobMaterial").
		Preload("Character.Talents.DungeonMaterial").
		Preload("Character.Talents.MobMaterial").
		Preload("Weapons").
		Preload("Echoes").
		Preload("Stats").
		Where("character_id = ?", id).
		First(&build).Error

	return build, err
}

// Создание нового репозитория
func NewWuwaRepository(db *gorm.DB) interfaces.RepoInterface[wuwa.WuwaCharacter, wuwa.WuwaBuild] {
	return &WuwaRepository{
		db: db,
	}
}
