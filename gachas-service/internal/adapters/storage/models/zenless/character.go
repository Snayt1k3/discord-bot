package zenless

import "gorm.io/gorm"

type ZenlessCharacter struct {
	gorm.Model
	Name        string
	Specialty   string
	Rank        string
	Attribute   string
	Faction     string
	AscensionID uint
	Ascension   ZenlessMaterial `gorm:"foreignKey:AscensionID;references:ID"`
	NodesID     uint
	Nodes       ZenlessNodes `gorm:"foreignKey:NodesID;references:ID"`
}

type ZenlessNodes struct {
	gorm.Model
	HuntBossMaterial   string
	ExpertBossMaterial string
	ResourceID         uint
	Resource           ZenlessMaterial `gorm:"foreignKey:ResourceID;references:ID"`
}

type ZenlessMaterial struct {
	gorm.Model
	Common string
	Rare   string
	Epic   string
}

func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&ZenlessMaterial{},
		&ZenlessNodes{},
		&ZenlessCharacter{},
		&ZenlessBuild{},
		&ZenlessDiscsPreset{},
		&ZenlessDiscs{},
		&ZenlessStats{},
		&ZenlessWeapons{},
		&ZenlessBuildDisc{},
		&ZenlessBuildWeapon{},
	)
}
