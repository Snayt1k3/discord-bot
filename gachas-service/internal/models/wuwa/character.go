package wuwa

import "gorm.io/gorm"

type Resonator struct {
	gorm.Model
	Name      string     `json:"name" gorm:"unique;not null"`
	Attribute string     `json:"attribute" gorm:"not null"` // Aero, Electro, Glacio и т.д.
	Role      string     `json:"role" gorm:"not null"`      // DPS, Support, Tank
	Ascension []Material `json:"ascension" gorm:"many2many:resonator_ascension_materials;"`
	Skills    []Skill    `json:"skills" gorm:"foreignKey:ResonatorID"`
	Weapons   []Weapon   `json:"weapons" gorm:"many2many:resonator_weapons;"`
	Builds    []Build    `json:"builds" gorm:"foreignKey:ResonatorID"`
}

type Skill struct {
	gorm.Model
	ResonatorID uint   `json:"resonator_id" gorm:"not null"`
	Name        string `json:"name" gorm:"not null"`
	Type        string `json:"type"` // Базовая атака, навык, ульта и т.д.
	Description string `json:"description"`
}
