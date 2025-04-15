package zenless

import "gorm.io/gorm"

type Agent struct {
	gorm.Model
	Name      string     `json:"name" gorm:"unique;not null"`
	Faction   string     `json:"faction" gorm:"not null"`
	Attribute string     `json:"attribute" gorm:"not null"` // Физика, Огонь, Электро и т. д.
	Role      string     `json:"role" gorm:"not null"`      // DPS, Tank, Support
	Ascension []Material `json:"ascension" gorm:"many2many:agent_ascension_materials;"`
	Skills    []Skill    `json:"skills" gorm:"foreignKey:AgentID"`
	Weapons   []Weapon   `json:"weapons" gorm:"many2many:agent_weapons;"`
	Builds    []Build    `json:"builds" gorm:"foreignKey:AgentID"`
}

type Skill struct {
	gorm.Model
	AgentID     uint   `json:"agent_id" gorm:"not null"`
	Name        string `json:"name" gorm:"not null"`
	Type        string `json:"type"` // Базовая атака, спец-навык, ульта и т. д.
	Description string `json:"description"`
}
