package genshin

import (
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type GenshinCommonMaterials struct {
	gorm.Model
	Common   string `json:"common" gorm:"unique;not null"`
	Uncommon string `json:"uncommon" gorm:"unique;not null"`
	Rare     string `json:"rare" gorm:"unique;not null"`
}

type GenshinAscension struct {
	gorm.Model
	Gem            string `json:"gem" gorm:"not null"`
	LocalSpecialty string `json:"local_specialty" gorm:"not null"`
	BossDrops      string `json:"boss_drops" gorm:"not null"`
}

type GenshinBooks struct {
	gorm.Model
	Common   string         `json:"common" gorm:"unique;not null"`
	Uncommon string         `json:"uncommon" gorm:"unique;not null"`
	Rare     string         `json:"rare" gorm:"unique;not null"`
	Weekdays pq.StringArray `json:"weekdays" gorm:"type:text[]"`
}

type GenshinTalent struct {
	gorm.Model
	BossDrops      string       `json:"boss_drops" gorm:"not null"`
	BooksID        uint         `json:"books_id"`
	Books          GenshinBooks `json:"books" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	TalentPriority string       `json:"talent_priority" gorm:"not null"`
}
