package zenless

import (
	"gorm.io/gorm"
)

type ZenlessBuild struct {
	gorm.Model
	CharacterID uint
	Character   ZenlessCharacter     `gorm:"foreignKey:CharacterID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	Weapons     []ZenlessWeapons     `gorm:"many2many:zenless_build_weapon;"`
	Discs       []ZenlessDiscsPreset `gorm:"many2many:zenless_build_disc;"`
	StatsID     uint                 `gorm:"not null"`
	Stats       ZenlessStats         `gorm:"foreignKey:StatsID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
}

type ZenlessDiscsPreset struct {
	gorm.Model
	BuildID        uint         `gorm:"not null"`
	FourPieceSetID uint         `gorm:"not null"`
	FourPieceSet   ZenlessDiscs `gorm:"foreignKey:FourPieceSetID;references:ID"`
	TwoPieceSetID  uint         `gorm:"not null"`
	TwoPieceSet    ZenlessDiscs `gorm:"foreignKey:TwoPieceSetID;references:ID"`
}

type ZenlessDiscs struct {
	gorm.Model
	Name           string `gorm:"not null"`
	TwoPieceBonus  string `gorm:"not null"`
	FourPieceBonus string `gorm:"not null"`
}

type ZenlessStats struct {
	gorm.Model
	FourthDisc       string `gorm:"type:varchar(100);not null"`
	FifthDisc        string `gorm:"type:varchar(100);not null"`
	SixthDisc        string `gorm:"type:varchar(100);not null"`
	SubStatsPriority string `gorm:"type:varchar(200);not null"`
}

type ZenlessWeapons struct {
	gorm.Model
	BuildID  uint   `gorm:"not null"`
	Name     string `gorm:"not null"`
	BaseATK  int32  `gorm:"not null"`
	Rarity   string `gorm:"not null"`
	Type     string `gorm:"not null"`
	Passive  string `gorm:"not null"`
	SubStat  string `gorm:"not null"`
	SubValue int32  `gorm:"not null"`
}

type ZenlessBuildWeapon struct {
	BuildID  uint `gorm:"primaryKey"`
	WeaponID uint `gorm:"primaryKey"`
	Priority int
}

type ZenlessBuildDisc struct {
	BuildID uint `gorm:"primaryKey"`
	DiscID  uint `gorm:"primaryKey"`
	Priority int
}