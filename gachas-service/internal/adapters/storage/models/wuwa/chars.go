package wuwa

import (
	"gorm.io/gorm"
)

type Character struct {
	gorm.Model
	Name        string    `gorm:"type:varchar(100);not null"`
	Element     string    `gorm:"type:varchar(50);not null"`
	Rarity      int       `gorm:"not null"`
	AscensionID uint      `gorm:"not null"`
	Ascension   Ascension `gorm:"foreignKey:AscensionID;references:ID"`
	TalentsID   uint      `gorm:"not null"`
	Talents     Talent    `gorm:"foreignKey:TalentsID;references:ID"`
}

func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&Character{},
		&Weapon{},
		&Ascension{},
		&Talent{},
		&Resource{},
		&Build{},
		&Echoes{},
		&Stats{},
	)
}
