package wuwa

import (
	"gorm.io/gorm"
)

type WuwaWeapon struct {
	gorm.Model
	Name       string `gorm:"type:varchar(100);not null"`
	WeaponType string `gorm:"type:varchar(50);not null"`
	Rarity     int    `gorm:"not null"`
	Passive    string `gorm:"type:text"`
	BaseATK    int    `gorm:"not null"`
}
