package repos

import (
	"gachas-service/internal/adapters/storage/models/genshin"
	"gachas-service/internal/interfaces"

	"gorm.io/gorm"
)

type GenshinRepository struct {
	db *gorm.DB
}

func (r *GenshinRepository) GetCharacters() ([]genshin.GenshinCharacter, error) {
	var characters []genshin.GenshinCharacter
	err := r.db.Find(&characters).Error

	if err != nil {
		return characters, err
	}

	return characters, nil
}

func (r *GenshinRepository) GetCharacterByID(id string) (genshin.GenshinCharacter, error) {
	var character genshin.GenshinCharacter

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

func (r *GenshinRepository) GetCharacterBuild(id string) (genshin.GenshinBuild, error) {
	var build genshin.GenshinBuild
	err := r.db.Preload("Character").
		Preload("Artifacts").
		Preload("Weapons").
		Preload("Teams").
		Preload("Stats").
		Where("character_id = ?", id).
		First(&build).Error

	if err != nil {
		return build, err
	}

	return build, nil
}

func NewGenshinRepository(db *gorm.DB) interfaces.RepoInterface[genshin.GenshinCharacter, genshin.GenshinBuild] {
	return &GenshinRepository{
		db: db,
	}
}
