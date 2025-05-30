package repos

import (
	"gachas-service/internal/adapters/storage/models/honkai"
	"gachas-service/internal/interfaces"
	"gorm.io/gorm"
)

type HonkaiRepository struct {
	db *gorm.DB
}

func (r *HonkaiRepository) GetCharacters() ([]honkai.HonkaiCharacter, error) {
	var characters []honkai.HonkaiCharacter
	err := r.db.Find(&characters).Error
	return characters, err
}

func (r *HonkaiRepository) GetCharacterByID(id string) (honkai.HonkaiCharacter, error) {
	var character honkai.HonkaiCharacter
	err := r.db.
		Preload("Ascension.Resource").
		Preload("Traces.DungeonResource").
		Preload("Traces.MobResource").
		First(&character, "id = ?", id).Error
	return character, err
}

func (r *HonkaiRepository) GetCharacterBuild(id string) (honkai.HonkaiBuild, error) {
	var build honkai.HonkaiBuild
	err := r.db.
		Preload("Character").
		Preload("Character.Ascension.Resource").
		Preload("Character.Traces.DungeonResource").
		Preload("Character.Traces.MobResource").
		Preload("Cones", func(db *gorm.DB) *gorm.DB {
			return db.
				Joins("JOIN honkai_build_cones ON honkai_light_cones.id = honkai_build_cones.cone_id").
				Order("honkai_build_cones.priority ASC")
		}).
		Preload("Artifacts", func(db *gorm.DB) *gorm.DB {
			return db.
				Joins("JOIN honkai_build_artifacts ON honkai_artifacts_presets.id = honkai_build_artifacts.artifact_id").
				Order("honkai_build_artifacts.priority ASC")
		}).
		Preload("Artifacts.Relics").
		Preload("Artifacts.Planars").
		Preload("Stats").
		Where("character_id = ?", id).
		First(&build).Error

	return build, err
}

func NewHonkaiRepository(db *gorm.DB) interfaces.RepoInterface[honkai.HonkaiCharacter, honkai.HonkaiBuild] {
	return &HonkaiRepository{
		db: db,
	}
}