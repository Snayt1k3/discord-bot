package repos

import (
	"gorm.io/gorm"
)

type GenshinRepository struct {
	db *gorm.DB
}


func NewGenshinRepository(db *gorm.DB) *GenshinRepository {
	return &GenshinRepository{
		db: db,
	}
}