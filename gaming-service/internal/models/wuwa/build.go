package wuwa

import "gorm.io/gorm"



type Weapon struct {
    gorm.Model
    Name     string `json:"name" gorm:"unique;not null"`
    Type     string `json:"type" gorm:"not null"`
    Rarity   int    `json:"rarity"`
    Passive  string `json:"passive"`
}

type Build struct {
    gorm.Model
    ResonatorID uint       `json:"resonator_id" gorm:"not null"`
    Name        string     `json:"name" gorm:"not null"`
    Echoes      []Echo     `json:"echoes" gorm:"many2many:build_echoes;"`
    Weapon      Weapon     `json:"weapon" gorm:"foreignKey:WeaponID"`
    WeaponID    uint       `json:"weapon_id"`
    Teams       []Team     `json:"teams" gorm:"many2many:build_teams;"`
}

type Echo struct {
    gorm.Model
    Name  string `json:"name" gorm:"unique;not null"`
    Set   string `json:"set"`
    Bonus string `json:"bonus"`
}

type Team struct {
    gorm.Model
    Name       string      `json:"name" gorm:"unique;not null"`
    Resonators []Resonator `json:"resonators" gorm:"many2many:team_resonators;"`
}
