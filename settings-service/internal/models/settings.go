package models

import (
	"gorm.io/gorm"
	"time"
)

type Settings struct {
	ID        uint            `gorm:"primaryKey"`
	GuildID   string          `gorm:"unique;not null"`
	Role      RolesSettings   `gorm:"foreignKey:GuildID;references:GuildID;constraint:OnDelete:CASCADE"`
	Welcome   WelcomeSettings `gorm:"foreignKey:GuildID;references:GuildID;constraint:OnDelete:CASCADE"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type RolesSettings struct {
	ID        uint   `gorm:"primaryKey"`
	GuildID   string `gorm:"not null;index"`
	MessageID string `json:"message_id"`
	Roles     []Role `gorm:"many2many:roles_settings_roles;constraint:OnDelete:CASCADE;"`
}

type Role struct {
	ID     uint   `gorm:"primaryKey"`
	RoleID string `json:"role_id"`
	Emoji  string `json:"emoji"`
}

type WelcomeSettings struct {
	ID        uint      `gorm:"primaryKey"`
	GuildID   string    `gorm:"not null;index"`
	ChannelId string    `json:"channel_id"`
	Messages  []Message `gorm:"many2many:welcome_setting_messages;constraint:OnDelete:CASCADE;"`
}

type Message struct {
	ID      uint   `gorm:"primaryKey"`
	Message string `json:"message"`
}

func Migrate(db *gorm.DB) {
	err := db.AutoMigrate(
		&Settings{},
		&RolesSettings{},
		&WelcomeSettings{},
		&Role{},
		&Message{},
	)
	if err != nil {
		panic("Migration Error: " + err.Error())
	}
}
