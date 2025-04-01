package zenless

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
    AgentID   uint       `json:"agent_id" gorm:"not null"`
    Name      string     `json:"name" gorm:"not null"`
    Discs     []Disc     `json:"discs" gorm:"many2many:build_discs;"`
    Weapon    Weapon     `json:"weapon" gorm:"foreignKey:WeaponID"`
    WeaponID  uint       `json:"weapon_id"`
    Teams     []Team     `json:"teams" gorm:"many2many:build_teams;"`
}

type Disc struct {
    gorm.Model
    Name  string `json:"name" gorm:"unique;not null"`
    Set   string `json:"set"`
    Bonus string `json:"bonus"`
}

type Team struct {
    gorm.Model
    Name     string   `json:"name" gorm:"unique;not null"`
    Agents   []Agent  `json:"agents" gorm:"many2many:team_agents;"`
}
