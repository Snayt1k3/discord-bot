package models

import "gorm.io/gorm"

type AutoModeSettings struct {
	gorm.Model
	GuildID     string            `gorm:"not null;index"`
	CapsLocks   []AntiCapsChannel `gorm:"many2many:auto_mode_capslock;constraint:OnDelete:CASCADE;"`
	AntiLinks   []AntiLinkChannel `gorm:"many2many:auto_mode_antilink;constraint:OnDelete:CASCADE;"`
	BannedWords []BannedWord      `gorm:"many2many:auto_mode_bannedwords;constraint:OnDelete:CASCADE;"`
	Enabled     bool              `gorm:"default:true"`
}

type BannedWord struct {
	ID   uint   `gorm:"primaryKey"`
	Word string `json:"word"`
}

type AntiCapsChannel struct {
	ID        uint   `gorm:"primaryKey"`
	ChannelID string `json:"channel_id"`
}

type AntiLinkChannel struct {
	ID        uint   `gorm:"primaryKey"`
	ChannelID string `json:"channel_id"`
}
