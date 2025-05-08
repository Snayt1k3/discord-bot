package dto

type GenshinCharacterBrief struct {
	ID      uint   `json:"id"`
	Name    string `json:"name"`
	Element string `json:"element"`
	Region  string `json:"region"`
}

type GenshinCommonMaterials struct {
	ID       uint   `json:"id"`
	Common   string `json:"common"`
	Uncommon string `json:"uncommon"`
	Rare     string `json:"rare"`
}

type GenshinAscensionMaterials struct {
	ID             uint   `json:"id"`
	LocalSpecialty string `json:"local_specialty"`
	Gem            string `json:"gem"`
	BossDrops      string `json:"boss_drops"`
}

type GenshinBooks struct {
	ID       uint     `json:"id"`
	Common   string   `json:"common"`
	Uncommon string   `json:"uncommon"`
	Rare     string   `json:"rare"`
	Weekdays []string `json:"weekdays"`
}

type GenshinTalentMaterials struct {
	ID             uint         `json:"id"`
	BossDrops      string       `json:"boss_drops"`
	Books          GenshinBooks `json:"books"`
	TalentPriority string       `json:"talent_priority"`
}

type GenshinCharacter struct {
	ID              uint                      `json:"id"`
	Name            string                    `json:"name"`
	WeaponType      string                    `json:"weapon_type"`
	Element         string                    `json:"element"`
	Region          string                    `json:"region"`
	Rarity          int                       `json:"rarity"`
	BaseStat        string                    `json:"base_stat"`
	Ascension       GenshinAscensionMaterials `json:"ascension"`
	Talents         GenshinTalentMaterials    `json:"talents"`
	CommonMaterials GenshinCommonMaterials    `json:"common_materials"`
}

type GenshinBuild struct {
	ID        uint              `json:"id"`
	Character GenshinCharacter  `json:"character"`
	Teams     []GenshinTeam     `json:"teams"`
	Artifacts []GenshinArtifact `json:"artifacts"`
	Weapons   []GenshinWeapon   `json:"weapons"`
	Stats     GenshinStats      `json:"stats"`
}

type GenshinArtifact struct {
	ID             uint   `json:"id"`
	Name           string `json:"name"`
	TwoPieceBonus  string `json:"two_piece_bonus"`
	FourPieceBonus string `json:"four_piece_bonus"`
}

type GenshinWeapon struct {
	ID       uint    `json:"id"`
	Name     string  `json:"name"`
	Type     string  `json:"type"`
	Rarity   int     `json:"rarity"`
	BaseATK  int     `json:"base_atk"`
	SubStat  string  `json:"sub_stat"`
	SubValue float64 `json:"sub_value"`
	Passive  string  `json:"passive"`
}

type GenshinStats struct {
	ID               uint   `json:"id"`
	Sands            string `json:"sands"`
	Goblet           string `json:"goblet"`
	Circlet          string `json:"circlet"`
	BestStats        string `json:"best_stats"`
	SubStatsPriority string `json:"sub_stats_priority"`
}

type GenshinTeam struct {
	ID         uint               `json:"id"`
	Characters []GenshinCharacter `json:"characters"`
}
