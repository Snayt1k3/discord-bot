package models

import "gorm.io/gorm"

type RolesSettings struct {
	gorm.Model
	GuildID   string `gorm:"not null;uniqueIndex"`
	MessageID string `json:"message_id"`
	Roles     []Role `gorm:"many2many:roles_settings_roles;constraint:OnDelete:CASCADE;"`
}

type Role struct {
	ID     uint   `gorm:"primaryKey"`
	RoleID string `json:"role_id"`
	Emoji  string `json:"emoji"`
}
