package honkai

import (
	"gorm.io/gorm"
)

type HonkaiBuild struct {
	gorm.Model
	CharacterID uint
	Character   HonkaiCharacter         `gorm:"foreignKey:CharacterID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	Cones       []HonkaiLightCones      `gorm:"many2many:honkai_build_cones;joinForeignKey:BuildID;joinReferences:ConeID"`
	Artifacts   []HonkaiArtifactsPreset `gorm:"many2many:honkai_build_artifacts;joinForeignKey:BuildID;joinReferences:ArtifactID"`
	StatsID     uint
	Stats       HonkaiStats `gorm:"foreignKey:StatsID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
}

type HonkaiLightCones struct {
	gorm.Model
	Name    string
	Rarity  int
	Path    string
	BaseDEF int32
	BaseHP  int32
	BaseATK int32
	Passive string
}

type HonkaiStats struct {
	gorm.Model
	Body         string
	Feet         string
	PlanarSphere string
	LinkRope     string
	SubStats     string
}

type HonkaiArtifactsPreset struct {
	gorm.Model
	RelicsID uint
	Relics   HonkaiRelics `gorm:"foreignKey:RelicsID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	PlanarID uint
	Planars  HonkaiPlanar `gorm:"foreignKey:PlanarID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
}

type HonkaiRelics struct {
	gorm.Model
	Name           string
	TwoPieceBonus  string
	FourPieceBonus string
}

type HonkaiPlanar struct {
	gorm.Model
	Name     string
	SetBonus string
}

type HonkaiBuildCone struct {
	BuildID  uint
	ConeID   uint
	Priority int
}

type HonkaiBuildArtifact struct {
	BuildID    uint
	ArtifactID uint
	Priority   int
}
