package dto

type GenshinCharacterBrief struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	WeaponType string `json:"weapon_type"`
	Element    string `json:"element"`
	Region     string `json:"region"`
	Rarity     int    `json:"rarity"`
}

type GenshinCommonMaterials struct {
	ID       string `json:"id"`
	Common   string `json:"common"`
	Uncommon string `json:"uncommon"`
	Rare     string `json:"rare"`
}

type GenshinAscensionMaterials struct {
	ID             string `json:"id"`
	LocalSpecialty string `json:"local_specialty"`
	BossDrops      string `json:"boss_drops"`
}

type GenshinBooks struct {
	ID       string `json:"id"`
	Common   string `json:"common"`
	Uncommon string `json:"uncommon"`
	Rare     string `json:"rare"`
}

type GenshinTalentMaterials struct {
	ID             string       `json:"id"`
	BossDrops      string       `json:"boss_drops"`
	Books          GenshinBooks `json:"books"`
	Weekdays       []string     `json:"weekdays"`
	TalentPriority string       `json:"talent_priority"`
}

type GenshinCharacter struct {
	ID              string                    `json:"id"`
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
	ID        string            `json:"id"`
	Character GenshinCharacter  `json:"character"`
	Teams     []GenshinTeam     `json:"teams"`
	Artifacts []GenshinArtifact `json:"artifacts"`
	Weapons   []GenshinWeapon   `json:"weapons"`
	Stats     GenshinStats      `json:"stats"`
}

type GenshinArtifact struct {
	ID             string `json:"id"`
	Name           string `json:"name"`
	Set            string `json:"set"`
	TwoPieceBonus  string `json:"two_piece_bonus"`
	FourPieceBonus string `json:"four_piece_bonus"`
}

type GenshinWeapon struct {
	ID       string  `json:"id"`
	Name     string  `json:"name"`
	Type     string  `json:"type"`
	Rarity   int     `json:"rarity"`
	BaseATK  int     `json:"base_atk"`
	SubStat  string  `json:"sub_stat"`
	SubValue float64 `json:"sub_value"`
	Passive  string  `json:"passive"`
}

type GenshinStats struct {
	ID               string `json:"id"`
	Sands            string `json:"sands"`
	Goblet           string `json:"goblet"`
	Circlet          string `json:"circlet"`
	BestStats        string `json:"best_stats"`
	SubStatsPriority string `json:"sub_stats_priority"`
}

type GenshinTeam struct {
	ID         string             `json:"id"`
	Characters []GenshinCharacter `json:"characters"`
}
