package wuwa

import (
	"gorm.io/gorm"
)

type WuwaBuild struct {
	gorm.Model
	CharacterID uint          `gorm:"not null"`
	Character   WuwaCharacter `gorm:"foreignKey:CharacterID;references:ID"`
	Weapons     []WuwaWeapon  `gorm:"many2many:build_weapons;"`
	Echoes      []WuwaEchoes  `gorm:"many2many:build_artifacts;"`
	StatsID     uint          `gorm:"not null"`
	Stats       WuwaStats     `gorm:"foreignKey:StatsID;references:ID"`
}

type WuwaEchoes struct {
	gorm.Model
	Name          string `gorm:"type:varchar(100);not null"`
	TwoPieceBonus string `gorm:"type:varchar(100);not null"`
	FullSetBonus  string `gorm:"type:varchar(100);not null"`
}

type WuwaStats struct {
	gorm.Model
	FourCostEchoStat  string `gorm:"type:varchar(100);not null"`
	ThreeCostEchoStat string `gorm:"type:varchar(100);not null"`
	SubStatsPriority  string `gorm:"type:varchar(100);not null"`
}
