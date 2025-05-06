package genshin

import (
	"gorm.io/gorm"
)

type Build struct {
	gorm.Model
	CharacterID uint       `json:"character_id" gorm:"not null"`
	Character   Character  `json:"character" gorm:"foreignKey:CharacterID"`
	Name        string     `json:"name" gorm:"not null"`
	Artifacts   []Artifact `json:"artifacts" gorm:"many2many:build_artifacts;"`
	Weapons     []Weapon   `json:"weapons" gorm:"many2many:build_weapons;"`
	Teams       []Team     `json:"teams" gorm:"many2many:build_teams;"`
	StatsId     uint       `json:"stats_id" gorm:"not null"`
	Stats       Stats      `json:"stats" gorm:"foreignKey:StatsId"`
}

type Artifact struct {
	gorm.Model
	Name           string `json:"name" gorm:"unique;not null"`
	TwoPieceBonus  string `json:"two_piece_bonus"`
	FourPieceBonus string `json:"four_piece_bonus"`
}

type Weapon struct {
	gorm.Model
	Name     string  `json:"name" gorm:"unique;not null"`
	Type     string  `json:"type"`
	Rarity   int     `json:"rarity"`
	BaseATK  int     `json:"base_atk"`
	SubStat  string  `json:"sub_stat"`
	SubValue float64 `json:"sub_value"`
	Passive  string  `json:"passive"`
}

type Team struct {
	gorm.Model
	Characters []Character `json:"characters" gorm:"many2many:team_characters;"`
}

type Stats struct {
	gorm.Model
	Sands            string `json:"sands" gorm:"not null"`
	Goblet           string `json:"goblet" gorm:"not null"`
	Circlet          string `json:"circlet" gorm:"not null"`
	BestStats        string `json:"best_stats" gorm:"not null"` // Todo: delete this field
	SubStatsPriority string `json:"sub_stats_priority" gorm:"not null"`
}

type BuildWeapon struct {
	BuildID  uint `gorm:"primaryKey"`
	WeaponID uint `gorm:"primaryKey"`
	Priority int
}

type BuildArtifact struct {
	BuildID    uint `gorm:"primaryKey"`
	ArtifactID uint `gorm:"primaryKey"`
	Priority   int
	SetGroup   int `gorm:"default:2"`
}
