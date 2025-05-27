package wuwa

import (
	"gorm.io/gorm"
)

type WuwaCharacter struct {
	gorm.Model
	Name        string        `gorm:"type:varchar(100);not null"`
	Element     string        `gorm:"type:varchar(50);not null"`
	WeaponType  string        `gorm:"type:varchar(50);not null"`
	Rarity      int           `gorm:"not null"`
	AscensionID uint          `gorm:"not null"`
	Ascension   WuwaAscension `gorm:"foreignKey:AscensionID;references:ID"`
	TalentsID   uint          `gorm:"not null"`
	Talents     WuwaTalent    `gorm:"foreignKey:TalentsID;references:ID"`
}

type WuwaAscension struct {
	gorm.Model
	LocalSpecialty string       `gorm:"type:varchar(100);not null"`
	BossMaterial   string       `gorm:"type:varchar(100);not null"`
	MobMaterialID  uint         `gorm:"not null"`
	MobMaterial    WuwaResource `gorm:"foreignKey:MobMaterialID;references:ID"`
}

type WuwaTalent struct {
	gorm.Model
	DungeonMaterial   WuwaResource `gorm:"foreignKey:DungeonMaterialID;references:ID"`
	DungeonMaterialID uint         `gorm:"not null"`
	MobMaterialID     uint         `gorm:"not null"`
	MobMaterial       WuwaResource `gorm:"foreignKey:MobMaterialID;references:ID"`
	BossMaterial      string       `gorm:"not null"`
}

type WuwaResource struct {
	gorm.Model
	UncommonName  string `gorm:"type:varchar(100);not null"`
	RareName      string `gorm:"type:varchar(100);not null"`
	EpicName      string `gorm:"type:varchar(100);not null"`
	LegendaryName string `gorm:"type:varchar(100);not null"`
}

func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&WuwaCharacter{},
		&WuwaWeapon{},
		&WuwaAscension{},
		&WuwaTalent{},
		&WuwaResource{},
		&WuwaBuild{},
		&WuwaEchoes{},
		&WuwaStats{},
	)
}
