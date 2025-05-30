package repos

import (
	"gachas-service/internal/adapters/storage/models/zenless"
	"gachas-service/internal/interfaces"
	"gorm.io/gorm"
)

type ZenlessRepository struct {
	db *gorm.DB
}

func (r *ZenlessRepository) GetCharacters() ([]zenless.ZenlessCharacter, error) {
	var characters []zenless.ZenlessCharacter
	err := r.db.Find(&characters).Error
	return characters, err
}

func (r *ZenlessRepository) GetCharacterByID(id string) (zenless.ZenlessCharacter, error) {
	var character zenless.ZenlessCharacter
	err := r.db.
		Preload("Ascension").
		Preload("Nodes.Resource").
		First(&character, "id = ?", id).Error
	return character, err
}

func (r *ZenlessRepository) GetCharacterBuild(id string) (zenless.ZenlessBuild, error) {
	var build zenless.ZenlessBuild
	err := r.db.
		Preload("Character").
		Preload("Character.Ascension").
		Preload("Character.Nodes.Resource").
		Preload("Discs.FourPieceSet").
		Preload("Discs.TwoPieceSet").
		Preload("Stats").
		Preload("Weapons", func(db *gorm.DB) *gorm.DB {
			return db.
				Joins("JOIN zenless_build_weapons ON zenless_weapons.id = zenless_build_weapons.weapon_id").
				Order("zenless_build_weapons.priority ASC")
		}).
		Preload("Discs", func(db *gorm.DB) *gorm.DB {
			return db.
				Joins("JOIN zenless_build_discs ON zenless_discs.id = zenless_build_discs.disc_id").
				Order("zenless_build_discs.priority ASC")
		}).
		Where("character_id = ?", id).
		First(&build).Error
	return build, err
}

func NewZenlessRepository(db *gorm.DB) interfaces.RepoInterface[zenless.ZenlessCharacter, zenless.ZenlessBuild] {
	return &ZenlessRepository{
		db: db,
	}
}
