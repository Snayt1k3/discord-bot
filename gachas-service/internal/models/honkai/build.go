package honkai

import "gorm.io/gorm"

type LightCone struct {
	gorm.Model
	Name    string `json:"name" gorm:"unique;not null"`
	Path    string `json:"path" gorm:"not null"`
	Rarity  int    `json:"rarity"`
	Passive string `json:"passive"`
}

type Build struct {
	gorm.Model
	CharacterID uint      `json:"character_id" gorm:"not null"`
	Name        string    `json:"name" gorm:"not null"`
	Relics      []Relic   `json:"relics" gorm:"many2many:build_relics;"`
	LightCone   LightCone `json:"light_cone" gorm:"foreignKey:LightConeID"`
	LightConeID uint      `json:"light_cone_id"`
	Teams       []Team    `json:"teams" gorm:"many2many:build_teams;"`
}

type Relic struct {
	gorm.Model
	Name  string `json:"name" gorm:"unique;not null"`
	Set   string `json:"set"`
	Bonus string `json:"bonus"`
}

type Team struct {
	gorm.Model
	Name       string      `json:"name" gorm:"unique;not null"`
	Characters []Character `json:"characters" gorm:"many2many:team_characters;"`
}
