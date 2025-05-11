package genshin

import "gorm.io/gorm"

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
