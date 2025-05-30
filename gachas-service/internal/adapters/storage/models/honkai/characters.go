package honkai

import (
	"gorm.io/gorm"
)

type HonkaiCharacter struct {
	gorm.Model
	Name        string
	Rarity      int
	Path        string
	Attribute   string
	AscensionID uint
	Ascension   HonkaiAscension `gorm:"foreignKey:AscensionID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	TracesID    uint
	Traces      HonkaiTraces    `gorm:"foreignKey:TracesID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
}

type HonkaiAscension struct {
	gorm.Model
	BossMaterial string
	ResourceID   uint
	Resource     HonkaiResource `gorm:"foreignKey:ResourceID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
}

type HonkaiResource struct {
	gorm.Model
	Uncommon string
	Rare     string
	Epic     string
}

type HonkaiTraces struct {
	gorm.Model
	DungeonResourceID uint
	DungeonResource   HonkaiResource `gorm:"foreignKey:DungeonResourceID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	MobResourceID     uint
	MobResource       HonkaiResource `gorm:"foreignKey:MobResourceID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	BossMaterial      string
}

func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&HonkaiCharacter{},
		&HonkaiAscension{},
		&HonkaiResource{},
		&HonkaiTraces{},
		&HonkaiBuild{},
		&HonkaiLightCones{},
		&HonkaiStats{},
		&HonkaiArtifactsPreset{},
		&HonkaiRelics{},
		&HonkaiPlanar{},
		&HonkaiBuildCone{},
		&HonkaiBuildArtifact{},
	)
}