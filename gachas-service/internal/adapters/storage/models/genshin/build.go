package genshin

import (
	"gorm.io/gorm"
)

type GenshinBuild struct {
	gorm.Model
	CharacterID uint              `json:"character_id" gorm:"not null"`
	Character   GenshinCharacter  `json:"character" gorm:"foreignKey:CharacterID"`
	Name        string            `json:"name" gorm:"not null"`
	Artifacts   []GenshinArtifact `json:"artifacts" gorm:"many2many:genshin_build_artifacts;"`
	Weapons     []GenshinWeapon   `json:"weapons" gorm:"many2many:genshin_build_weapons;"`
	Teams       []GenshinTeam     `json:"teams" gorm:"many2many:genshin_build_teams;"`
	StatsId     uint              `json:"stats_id" gorm:"not null"`
	Stats       GenshinStats      `json:"stats" gorm:"foreignKey:StatsId"`
}

type GenshinArtifact struct {
	gorm.Model
	Name           string `json:"name" gorm:"unique;not null"`
	TwoPieceBonus  string `json:"two_piece_bonus" gorm:"type:text"`
	FourPieceBonus string `json:"four_piece_bonus" gorm:"type:text"`
}

type GenshinWeapon struct {
	gorm.Model
	Name     string  `json:"name" gorm:"unique;not null"`
	Type     string  `json:"type"`
	Rarity   int     `json:"rarity"`
	BaseATK  int     `json:"base_atk"`
	SubStat  string  `json:"sub_stat"`
	SubValue float64 `json:"sub_value"`
	Passive  string  `json:"passive"`
}

type GenshinTeam struct {
	gorm.Model
	Characters []GenshinCharacter `json:"characters" gorm:"many2many:team_characters;"`
}

type GenshinStats struct {
	gorm.Model
	Sands            string `json:"sands" gorm:"not null"`
	Goblet           string `json:"goblet" gorm:"not null"`
	Circlet          string `json:"circlet" gorm:"not null"`
	SubStatsPriority string `json:"sub_stats_priority" gorm:"not null"`
}

type GenshinBuildWeapon struct {
	BuildID  uint `gorm:"primaryKey"`
	WeaponID uint `gorm:"primaryKey"`
	Priority int
}

type GenshinBuildArtifact struct {
	BuildID    uint `gorm:"primaryKey"`
	ArtifactID uint `gorm:"primaryKey"`
	Priority   int
	SetGroup   int `gorm:"default:2"`
}
