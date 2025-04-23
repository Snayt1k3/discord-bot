package genshin

import (
	"gorm.io/gorm"
	"github.com/lib/pq"
)

type CommonMaterials struct {
	gorm.Model
	Common   string `json:"common" gorm:"unique;not null"`
	Uncommon string `json:"uncommon" gorm:"unique;not null"`
	Rare     string `json:"rare" gorm:"unique;not null"`
}

type AscensionMaterials struct {
	gorm.Model
	LocalSpecialty string `json:"local_specialty" gorm:"not null"`
	BossDrops      string `json:"boss_drops" gorm:"not null"`
}

type Books struct {
	gorm.Model
	Common   string `json:"common" gorm:"unique;not null"`
	Uncommon string `json:"uncommon" gorm:"unique;not null"`
	Rare     string `json:"rare" gorm:"unique;not null"`
	Weekdays       pq.StringArray `json:"weekdays" gorm:"type:text[]"`
}

type TalentMaterials struct {
	gorm.Model
	BossDrops      string   `json:"boss_drops" gorm:"not null"`
	BooksID        uint     `json:"books_id"`
	Books          Books    `json:"books" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	TalentPriority string   `json:"talent_priority" gorm:"not null"`
}
