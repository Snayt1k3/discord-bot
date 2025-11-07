package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID            uint      `gorm:"primaryKey"`
	UserID        string    `gorm:"not null;index:idx_user_guild,unique"`
	GuildID       string    `gorm:"not null;index:idx_user_guild,unique"`
	Experience    int       `json:"experience"`
	Level         int       `json:"level"`
	LastMessageAt time.Time `json:"last_message_at"`
	NextLevelXP   int       `json:"next_level_xp"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt `gorm:"index"`
}

func Migrate(db *gorm.DB) {
	err := db.AutoMigrate(
		&User{},
	)
	if err != nil {
		panic("Migration Error: " + err.Error())
	}
}
