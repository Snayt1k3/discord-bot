package models

import "gorm.io/gorm"

type WelcomeSettings struct {
	gorm.Model
	GuildID   string    `gorm:"not null;uniqueIndex"`
	ChannelId string    `json:"channel_id"`
	Messages  []Message `gorm:"many2many:welcome_setting_messages;constraint:OnDelete:CASCADE;"`
}

type Message struct {
	ID      uint   `gorm:"primaryKey"`
	Message string `json:"message"`
}
