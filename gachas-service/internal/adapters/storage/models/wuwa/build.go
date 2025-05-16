package wuwa

import (
	"gorm.io/gorm"
)

type WuwaBuild struct {
	gorm.Model
	CharacterID     uint          `gorm:"not null"`
	Character       WuwaCharacter `gorm:"foreignKey:CharacterID;references:ID"`
	Weapons         []WuwaWeapon  `gorm:"many2many:build_wuwa_weapons;"`
	Echoes          []WuwaEchoes  `gorm:"many2many:build_wuwa_echoes;"`
	StatsID         uint          `gorm:"not null"`
	Stats           WuwaStats     `gorm:"foreignKey:StatsID;references:ID"`
	BestPrimaryEcho string        `gorm:"type:varchar(100);not null"`
}

type WuwaEchoes struct {
	gorm.Model
	Name          string `gorm:"type:varchar(100);not null"`
	TwoPieceBonus string `gorm:"type:varchar(300);not null"`
	FullSetBonus  string `gorm:"type:varchar(300);not null"`
}

type WuwaStats struct {
	gorm.Model
	FourCostEchoStat  string `gorm:"type:varchar(100);not null"`
	ThreeCostEchoStat string `gorm:"type:varchar(100);not null"`
	OneCostEchoStat   string `gorm:"type:varchar(100);not null"`
	SubStatsPriority  string `gorm:"type:varchar(200);not null"`
}

type WuwaWeapon struct {
	gorm.Model
	Name       string  `gorm:"type:varchar(50);not null"`
	WeaponType string  `gorm:"type:varchar(50);not null"`
	Rarity     int     `gorm:"not null"`
	Passive    string  `gorm:"type:text"`
	BaseATK    int     `gorm:"not null"`
	SubStat    string  `gorm:"type:varchar(50);not null"`
	SubValue   float32 `gorm:"not null"`
}
