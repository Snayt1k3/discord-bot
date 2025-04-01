package genshin

import "gorm.io/gorm"


type Build struct {
    gorm.Model
    CharacterID uint        `json:"character_id" gorm:"not null"`
    Name        string      `json:"name" gorm:"not null"`
    Artifacts   []Artifact  `json:"artifacts" gorm:"many2many:build_artifacts;"`
    Weapon      Weapon      `json:"weapon" gorm:"foreignKey:WeaponID"`
    WeaponID    uint        `json:"weapon_id"`
    Teams       []Character `json:"teams" gorm:"many2many:build_teams;"`
}

type Artifact struct {
    gorm.Model
    Name  string `json:"name" gorm:"unique;not null"`
    Set   string `json:"set"`
    Bonus string `json:"bonus"`
}

type Weapon struct {
    gorm.Model
    Name     string `json:"name" gorm:"unique;not null"`
    Type     string `json:"type"`
    Rarity   int    `json:"rarity"`
    Passive  string `json:"passive"`
}

type Team struct {
    gorm.Model
    Name       string      `json:"name" gorm:"unique;not null"`
    Characters []Character `json:"characters" gorm:"many2many:team_characters;"`
}