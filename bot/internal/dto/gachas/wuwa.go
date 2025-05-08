package dto

type WuwaCharacterShort struct {
	ID      uint   `json:"id"`
	Name    string `json:"name"`
	Element string `json:"element"`
	Rarity  int    `json:"rarity"`
}

type WuwaResource struct {
	UncommonName  string `json:"uncommon_name"`
	RareName      string `json:"rare_name"`
	EpicName      string `json:"epic_name"`
	LegendaryName string `json:"legendary_name"`
}

type WuwaAscension struct {
	LocalSpecialty string       `json:"local_specialty"`
	BossMaterial   string       `json:"boss_material"`
	MobMaterial    WuwaResource `json:"mob_material"`
}

type WuwaTalent struct {
	DungeonMaterial WuwaResource `json:"dungeon_material"`
	MobMaterial     WuwaResource `json:"mob_material"`
	BossMaterial    string
}

type WuwaWeapon struct {
	Name       string  `json:"name"`
	WeaponType string  `json:"weapon_type"`
	Rarity     int     `json:"rarity"`
	Passive    string  `json:"passive"`
	BaseATK    int     `json:"base_atk"`
	SubStat    string  `json:"sub_stat"`
	SubValue   float32 `json:"sub_value"`
}

type WuwaCharacterFull struct {
	ID          uint          `json:"id"`
	Name        string        `json:"name"`
	WeaponType  string        `json:"weapon_type"`
	Affiliation string        `json:"affiliation"`
	Element     string        `json:"element"`
	Rarity      int           `json:"rarity"`
	Ascension   WuwaAscension `json:"ascension"`
	Talents     WuwaTalent    `json:"talents"`
}

type WuwaEcho struct {
	Name          string `json:"name"`
	TwoPieceBonus string `json:"two_piece_bonus"`
	FullSetBonus  string `json:"full_set_bonus"`
}

type WuwaStats struct {
	FourCostEchoStat  string `json:"four_cost_echo_stat"`
	ThreeCostEchoStat string `json:"three_cost_echo_stat"`
	SubStatsPriority  string `json:"sub_stats_priority"`
}

type WuwaCharacterBuild struct {
	Character       WuwaCharacterFull `json:"character"`
	Weapons         []WuwaWeapon      `json:"weapons"`
	Echoes          []WuwaEcho        `json:"echoes"`
	Stats           WuwaStats         `json:"stats"`
	BestPrimaryEcho string            `json:"best_primary_echo"`
}
