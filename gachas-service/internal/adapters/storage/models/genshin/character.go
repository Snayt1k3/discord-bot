package genshin

import (
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type GenshinCharacter struct {
	gorm.Model
	Name              string                 `json:"name" gorm:"unique;not null"`
	Element           string                 `json:"element" gorm:"not null"`
	WeaponType        string                 `json:"weapon_type" gorm:"not null"`
	BaseStat          string                 `json:"base_stat" gorm:"not null"`
	Region            string                 `json:"region" gorm:"not null"`
	Rarity            int                    `json:"rarity" gorm:"not null"`
	AscensionID       uint                   `json:"ascension_id" gorm:"not null"`
	Ascension         GenshinAscension       `json:"ascension" gorm:"foreignKey:AscensionID"`
	TalentsID         uint                   `json:"talents_id" gorm:"not null"`
	Talents           GenshinTalent          `json:"talents" gorm:"foreignKey:TalentsID"`
	CommonMaterialsID uint                   `json:"common_materials_id" gorm:"not null"`
	CommonMaterials   GenshinCommonMaterials `json:"common_materials" gorm:"foreignKey:CommonMaterialsID"`
}

type GenshinCommonMaterials struct {
	gorm.Model
	Common   string `json:"common" gorm:"unique;not null"`
	Uncommon string `json:"uncommon" gorm:"unique;not null"`
	Rare     string `json:"rare" gorm:"unique;not null"`
}

type GenshinAscension struct {
	gorm.Model
	Gem            string `json:"gem" gorm:"not null"`
	LocalSpecialty string `json:"local_specialty" gorm:"not null"`
	BossDrops      string `json:"boss_drops" gorm:"not null"`
}

type GenshinBooks struct {
	gorm.Model
	Common   string         `json:"common" gorm:"unique;not null"`
	Uncommon string         `json:"uncommon" gorm:"unique;not null"`
	Rare     string         `json:"rare" gorm:"unique;not null"`
	Weekdays pq.StringArray `json:"weekdays" gorm:"type:text[]"`
}

type GenshinTalent struct {
	gorm.Model
	BossDrops      string       `json:"boss_drops" gorm:"not null"`
	BooksID        uint         `json:"books_id"`
	Books          GenshinBooks `json:"books" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	TalentPriority string       `json:"talent_priority" gorm:"not null"`
}

func Migrate(db *gorm.DB) error {
	if err := db.AutoMigrate(
		&GenshinCharacter{},
		&GenshinBuild{},
		&GenshinArtifact{},
		&GenshinWeapon{},
		&GenshinTeam{},
		&GenshinBuildWeapon{},
		&GenshinBuildArtifact{},
		&GenshinStats{},
		&GenshinCommonMaterials{},
		&GenshinAscension{},
		&GenshinTalent{},
		&GenshinBooks{},
	); err != nil {
		return err
	}
	return nil
}
