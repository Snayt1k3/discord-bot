package dto

type ZenlessCharacterDTO struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	Specialty string `json:"specialty"`
	Rank      string `json:"rank"`
	Attribute string `json:"attribute"`
	Faction   string `json:"faction"`
}

type ZenlessAscensionDTO struct {
	Common string `json:"common"`
	Rare   string `json:"rare"`
	Epic   string `json:"epic"`
}

type ZenlessNodesDTO struct {
	HuntBossMaterial   string            `json:"hunt_boss_material"`
	ExpertBossMaterial string            `json:"expert_boss_material"`
	Denny              int32             `json:"denny"`
	HamsterCagePass    int               `json:"hamster_cage_pass"`
	Resource           ZenlessAscensionDTO `json:"resource"`
}

type ZenlessCharacterFull struct {
	ID        uint               `json:"id"`
	Name      string             `json:"name"`
	Specialty string             `json:"specialty"`
	Rank      string             `json:"rank"`
	Attribute string             `json:"attribute"`
	Faction   string             `json:"faction"`
	Ascension ZenlessAscensionDTO `json:"ascension"`
	Nodes     ZenlessNodesDTO     `json:"nodes"`
}

type ZenlessWeaponDTO struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	BaseATK  int32  `json:"base_atk"`
	Rarity   string `json:"rarity"`
	Type     string `json:"type"`
	Passive  string `json:"passive"`
	SubStat  string `json:"sub_stat"`
	SubValue int32  `json:"sub_value"`
	Priority int    `json:"priority"`
}

type ZenlessDiscSetDTO struct {
	Name           string `json:"name"`
	TwoPieceBonus  string `json:"two_piece_bonus"`
	FourPieceBonus string `json:"four_piece_bonus"`
}

type ZenlessDiscDTO struct {
	TwoPiece ZenlessDiscSetDTO `json:"two_piece"`
	FourPiece ZenlessDiscSetDTO `json:"four_piece"`
	Priority int                `json:"priority"`
}

type ZenlessStatsDTO struct {
	FourthDisc       string `json:"fourth_disc"`
	FifthDisc        string `json:"fifth_disc"`
	SixthDisc        string `json:"sixth_disc"`
	SubStatsPriority string `json:"sub_stats_priority"`
}

type ZenlessBuildDTO struct {
	ID        uint                `json:"id"`
	Character ZenlessCharacterFull `json:"character"`
	Weapons   []ZenlessWeaponDTO `json:"weapons"`
	Discs     []ZenlessDiscDTO   `json:"discs"`
	Stats     ZenlessStatsDTO    `json:"stats"`
}