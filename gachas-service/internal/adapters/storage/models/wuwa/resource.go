package wuwa

import (
	"gorm.io/gorm"
)

type Ascension struct {
	gorm.Model
	LocalSpecialty string       `gorm:"type:varchar(100);not null"`
	BossMaterial   string       `gorm:"type:varchar(100);not null"`
	MobMaterialID  uint         `gorm:"not null"`
	MobMaterial    Resource `gorm:"foreignKey:MobMaterialID;references:ID"`
}

type Talent struct {
	gorm.Model
	DungeonMaterial   Resource `gorm:"foreignKey:DungeonMaterialID;references:ID"`
	DungeonMaterialID uint         `gorm:"not null"`
	MobMaterialID     uint         `gorm:"not null"`
	MobMaterial       Resource `gorm:"foreignKey:MobMaterialID;references:ID"`
	BossMaterial 	  string `gorm:"not null"`
}

type Resource struct {
	gorm.Model
	UncommonName  string `gorm:"type:varchar(100);not null"`
	RareName      string `gorm:"type:varchar(100);not null"`
	EpicName      string `gorm:"type:varchar(100);not null"`
	LegendaryName string `gorm:"type:varchar(100);not null"`
}
