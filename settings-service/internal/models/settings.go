package models

import (
	"gorm.io/gorm"
)

type Settings struct {
	gorm.Model
	GuildID  string           `gorm:"not null;uniqueIndex"`
	Role     RolesSettings    `gorm:"foreignKey:GuildID;references:GuildID;constraint:OnDelete:CASCADE"`
	Welcome  WelcomeSettings  `gorm:"foreignKey:GuildID;references:GuildID;constraint:OnDelete:CASCADE"`
	AutoMode AutoModeSettings `gorm:"foreignKey:GuildID;references:GuildID;constraint:OnDelete:CASCADE"`
	Log      LogSettings      `gorm:"foreignKey:GuildID;references:GuildID;constraint:OnDelete:CASCADE"`
}

func Migrate(db *gorm.DB) {
	err := db.AutoMigrate(
		&Settings{},
		&RolesSettings{},
		&WelcomeSettings{},
		&Role{},
		&Message{},
		&AutoModeSettings{},
		&BannedWord{},
		&AntiCapsChannel{},
		&AntiLinkChannel{},
		&LogSettings{},
		&LogEvent{},
	)
	if err != nil {
		panic("Migration Error: " + err.Error())
	}
}
