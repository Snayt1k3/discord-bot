package genshin

import "gorm.io/gorm"

type Character struct {
    gorm.Model
    Name        string     `json:"name" gorm:"unique;not null"`
    Element     string     `json:"element" gorm:"not null"`
    WeaponType  string     `json:"weapon_type" gorm:"not null"`
    Ascension   []Material `json:"ascension" gorm:"many2many:character_ascension_materials;"`
    TalentBooks []Material `json:"talent_books" gorm:"many2many:character_talent_books;"`
    Builds      []Build    `json:"builds" gorm:"foreignKey:CharacterID"`
}