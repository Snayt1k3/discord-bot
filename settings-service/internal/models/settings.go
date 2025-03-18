package models

import (
	"encoding/json"
	"time"

	"gorm.io/gorm"
)

type RoleSetting struct {
	ID        uint            `gorm:"primaryKey"`
	GuildID   string          `gorm:"not null;index"`
	MessageID string          `json:"message_id" gorm:"not null"`
	Role      json.RawMessage `gorm:"type:json"`
}

type WelcomeSetting struct {
	ID        uint            `gorm:"primaryKey"`
	GuildID   string          `gorm:"not null;index"`
	ChannelId string          `json:"channel_id" gorm:"not null"`
}

type GuildSetting struct {
	ID        uint           `gorm:"primaryKey"`
	GuildID   string         `gorm:"unique;not null"`
	Role      RoleSetting    `gorm:"foreignKey:GuildID;references:GuildID;constraint:OnDelete:CASCADE"` 
	Welcome   WelcomeSetting `gorm:"foreignKey:GuildID;references:GuildID;constraint:OnDelete:CASCADE"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func Migrate(db *gorm.DB) {
	err := db.AutoMigrate(&GuildSetting{}, &RoleSetting{}, &WelcomeSetting{})
	if err != nil {
		panic("Migration Error: " + err.Error())
	}
}
