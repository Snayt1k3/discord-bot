package models

import (
	"time"

	"gorm.io/gorm"
)


type GuildSetting struct {
	ID      uint   `gorm:"primaryKey"` 
	GuildID string `gorm:"unique"`     
	Settings SettingsJson `gorm:"type:jsonb;not null"`
	CreatedAt    time.Time      
	UpdatedAt    time.Time     
}

type BotSetting struct {
	ID           uint   `gorm:"primaryKey"`         
	BotStatus    string `gorm:"type:text;not null"` 
	Description  string `gorm:"type:text"`         
	HelpMessage  string `gorm:"type:text"`          
	HelloMessage string `gorm:"type:text"`
	CreatedAt    time.Time      
	UpdatedAt    time.Time    
}

func AutoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(&BotSetting{}, &GuildSetting{})
}