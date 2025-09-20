package models

import (
	"time"

	"gorm.io/gorm"
)

type Settings struct {
	ID        uint             `gorm:"primaryKey"`
	GuildID   string           `gorm:"not null;uniqueIndex"`
	Role      RolesSettings    `gorm:"foreignKey:GuildID;references:GuildID;constraint:OnDelete:CASCADE"`
	Welcome   WelcomeSettings  `gorm:"foreignKey:GuildID;references:GuildID;constraint:OnDelete:CASCADE"`
	AutoMode  AutoModeSettings `gorm:"foreignKey:GuildID;references:GuildID;constraint:OnDelete:CASCADE"`
	Log       LogSettings      `gorm:"foreignKey:GuildID;references:GuildID;constraint:OnDelete:CASCADE"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type RolesSettings struct {
	ID        uint   `gorm:"primaryKey"`
	GuildID   string `gorm:"not null;uniqueIndex"`
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
	GuildID   string    `gorm:"not null;uniqueIndex"`
	ChannelId string    `json:"channel_id"`
	Messages  []Message `gorm:"many2many:welcome_setting_messages;constraint:OnDelete:CASCADE;"`
}

type Message struct {
	ID      uint   `gorm:"primaryKey"`
	Message string `json:"message"`
}

type AutoModeSettings struct {
	ID          uint              `gorm:"primaryKey"`
	GuildID     string            `gorm:"not null;index"`
	CapsLocks   []AntiCapsChannel `gorm:"many2many:auto_mode_capslock;joinForeignKey:AutoModeID;joinReferences:ChannelID;constraint:OnDelete:CASCADE;"`
	AntiLinks   []AntiLinkChannel `gorm:"many2many:auto_mode_antilink;joinForeignKey:AutoModeID;joinReferences:ChannelID;constraint:OnDelete:CASCADE;"`
	BannedWords []BannedWord      `gorm:"many2many:auto_mode_bannedwords;joinForeignKey:AutoModeID;joinReferences:WordID;constraint:OnDelete:CASCADE;"`
	Enabled     bool              `json:"enabled"`
}

type BannedWord struct {
	ID      uint   `gorm:"primaryKey"`
	GuildID string `gorm:"not null;uniqueIndex"`
	Word    string `json:"word"`
}

type AntiCapsChannel struct {
	ID        uint   `gorm:"primaryKey"`
	GuildID   string `gorm:"not null;uniqueIndex"`
	ChannelID string `json:"channel_id"`
}

type AntiLinkChannel struct {
	ID        uint   `gorm:"primaryKey"`
	GuildID   string `gorm:"not null;uniqueIndex"`
	ChannelID string `json:"channel_id"`
}

type LogSettings struct {
	ID        uint   `gorm:"primaryKey"`
	GuildID   string `gorm:"not null;uniqueIndex"`
	ChannelID string `json:"channel_id"`
	Enabled   bool
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
	)
	if err != nil {
		panic("Migration Error: " + err.Error())
	}
}
