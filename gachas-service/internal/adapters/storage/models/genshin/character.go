package genshin

import "gorm.io/gorm"

type Character struct {
	gorm.Model
	Name              string             `json:"name" gorm:"unique;not null"`
	Element           string             `json:"element" gorm:"not null"`
	WeaponType        string             `json:"weapon_type" gorm:"not null"`
	BaseStat          string             `json:"base_stat" gorm:"not null"`
	Region            string             `json:"region" gorm:"not null"`
	Rarity            int                `json:"rarity" gorm:"not null"`
	AscensionID       uint               `json:"ascension_id" gorm:"not null"`
	Ascension         AscensionMaterials `json:"ascension" gorm:"foreignKey:AscensionID"`
	TalentsID         uint               `json:"talents_id" gorm:"not null"`
	Talents           TalentMaterials    `json:"talents" gorm:"foreignKey:TalentsID"`
	CommonMaterialsID uint               `json:"common_materials_id" gorm:"not null"`
	CommonMaterials   CommonMaterials    `json:"common_materials" gorm:"foreignKey:CommonMaterialsID"`
}

func Migrate(db *gorm.DB) error {
	if err := db.AutoMigrate(
		&Character{},
		&Build{},
		&Artifact{},
		&Weapon{},
		&Team{},
		&BuildWeapon{},
		&BuildArtifact{},
		&Stats{},
		&CommonMaterials{},
		&AscensionMaterials{},
		&TalentMaterials{},
		&Books{},
	); err != nil {
		return err
	}
	return nil
}
