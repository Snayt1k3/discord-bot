package models

import (
	"time"
	"settings-service/internal/dto"
	"gorm.io/gorm"
)


type GuildSetting struct {
	ID      uint   `gorm:"primaryKey"` 
	GuildID string `gorm:"unique"`     
	Roles dto.RolesSettings `gorm:"type:json;not null"`
	CreatedAt    time.Time      
	UpdatedAt    time.Time     
}

type BotSetting struct {
	ID           uint   `gorm:"primaryKey"`         
	BotStatus    string `gorm:"type:text;not null"` 
	Description  string `gorm:"type:text"`         
	HelpMessage  string `gorm:"type:text"`          
	HelloMessages []string  `gorm:"type:json"`
	CreatedAt    time.Time      
	UpdatedAt    time.Time    
}

func AutoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(&BotSetting{}, &GuildSetting{})
}