package honkai

import "gorm.io/gorm"

type Character struct {
	gorm.Model
	Name       string      `json:"name" gorm:"unique;not null"`
	Path       string      `json:"path" gorm:"not null"` // Тип пути (например, Охота, Разрушение)
	Element    string      `json:"element" gorm:"not null"`
	Ascension  []Material  `json:"ascension" gorm:"many2many:character_ascension_materials;"`
	Skills     []Skill     `json:"skills" gorm:"foreignKey:CharacterID"`
	LightCones []LightCone `json:"light_cones" gorm:"many2many:character_light_cones;"`
	Builds     []Build     `json:"builds" gorm:"foreignKey:CharacterID"`
}

type Skill struct {
	gorm.Model
	CharacterID uint   `json:"character_id" gorm:"not null"`
	Name        string `json:"name" gorm:"not null"`
	Type        string `json:"type"` // Базовая атака, навык, ульта, техника
	Description string `json:"description"`
}
