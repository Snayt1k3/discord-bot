package models

import (
	"time"
	"settings-service/internal/dto"
	"gorm.io/gorm"
)


type GuildSetting struct {
	ID      uint   `gorm:"primaryKey"` 
	GuildID string `gorm:"unique"`     
	Roles dto.RolesSettings `gorm:"type:json"`
	CreatedAt    time.Time      
	UpdatedAt    time.Time     
}

func AutoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(&GuildSetting{})
}