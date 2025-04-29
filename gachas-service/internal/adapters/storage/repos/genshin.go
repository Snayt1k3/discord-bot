package repos

import (
	"gachas-service/internal/adapters/storage/models/genshin"
	"gachas-service/internal/interfaces"
	"gorm.io/gorm"
)

type GenshinRepository struct {
	db *gorm.DB
}

func (r *GenshinRepository) GetCharacters() ([]genshin.Character, error) {
	var characters []genshin.Character
	err := r.db.Find(&characters).Error

	if err != nil {
		return characters, err
	}

	return characters, nil
}

func (r *GenshinRepository) GetCharacterByID(id string) (genshin.Character, error) {
	var character genshin.Character

	err := r.db.
		Preload("Ascension").
		Preload("Talents").
		Preload("Talents.Books").
		Preload("CommonMaterials").
		First(&character, "id = ?", id).Error

	if err != nil {
		return character, err
	}

	return character, nil
}

func (r *GenshinRepository) GetCharacterBuilds(id string) ([]genshin.Build, error) {
	var builds []genshin.Build
	err := r.db.Preload("Character"). 
		Preload("Artifacts").         
		Preload("Weapons").            
		Preload("Teams").              
		Preload("Stats").             
		Where("character_id = ?", id).
		Find(&builds).Error

	if err != nil {
		return builds, err
	}

	return builds, nil
}

func NewGenshinRepository(db *gorm.DB) interfaces.RepoInterface[genshin.Character, genshin.Build] {
	return &GenshinRepository{
		db: db,
	}
}
