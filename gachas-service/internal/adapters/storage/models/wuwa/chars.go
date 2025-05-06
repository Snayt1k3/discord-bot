package wuwa

import (
	"gorm.io/gorm"
)

type WuwaCharacter struct {
	gorm.Model
	Name        string        `gorm:"type:varchar(100);not null"`
	Element     string        `gorm:"type:varchar(50);not null"`
	WeaponId    uint          `gorm:"not null"`
	Weapon      WuwaWeapon    `gorm:"foreignKey:WeaponId;references:ID"`
	Rarity      int           `gorm:"not null"`
	AscensionID uint          `gorm:"not null"`
	Ascension   WuwaAscension `gorm:"foreignKey:AscensionID;references:ID"`
	TalentsID   uint          `gorm:"not null"`
	Talents     WuwaTalent    `gorm:"foreignKey:TalentsID;references:ID"`
}
