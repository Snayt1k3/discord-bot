package wuwa

import (
	"gorm.io/gorm"
)

type Build struct {
	gorm.Model
	CharacterID uint          `gorm:"not null"`
	Character   Character `gorm:"foreignKey:CharacterID;references:ID"`
	Weapons     []Weapon  `gorm:"many2many:build_wuwa_weapons;"`
	Echoes      []Echoes  `gorm:"many2many:build_wuwa_echoes;"`
	StatsID     uint          `gorm:"not null"`
	Stats       Stats     `gorm:"foreignKey:StatsID;references:ID"`
	BestPrimaryEcho string 	  `gorm:"type:varchar(100);not null"`
}

type Echoes struct {
	gorm.Model
	Name          string `gorm:"type:varchar(100);not null"`
	TwoPieceBonus string `gorm:"type:varchar(100);not null"`
	FullSetBonus  string `gorm:"type:varchar(100);not null"`
}

type Stats struct {
	gorm.Model
	FourCostEchoStat  string `gorm:"type:varchar(100);not null"`
	ThreeCostEchoStat string `gorm:"type:varchar(100);not null"`
	SubStatsPriority  string `gorm:"type:varchar(100);not null"`
}
