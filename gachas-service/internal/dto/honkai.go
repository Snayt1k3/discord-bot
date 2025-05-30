package dto

type HonkaiCharacterDTO struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	Rarity    int    `json:"rarity"`
	Path      string `json:"path"`
	Attribute string `json:"attribute"`
	Ascension HonkaiAscensionDTO `json:"ascension"`
	Traces    HonkaiTracesDTO    `json:"traces"`
}

type HonkaiCharacterBriefDTO struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	Rarity    int    `json:"rarity"`
	Path      string `json:"path"`
	Attribute string `json:"attribute"`
}

type HonkaiAscensionDTO struct {
	ID            uint               `json:"id"`
	BossMaterial  string             `json:"boss_material"`
	Resource      HonkaiResourceDTO  `json:"resource"`
}

type HonkaiResourceDTO struct {
	ID       uint   `json:"id"`
	Uncommon string `json:"uncommon"`
	Rare     string `json:"rare"`
	Epic     string `json:"epic"`
}

type HonkaiTracesDTO struct {
	ID              uint              `json:"id"`
	DungeonResource HonkaiResourceDTO `json:"dungeon_resource"`
	MobResource     HonkaiResourceDTO `json:"mob_resource"`
	BossMaterial    string            `json:"boss_material"`
}

type HonkaiBuildDTO struct {
	ID         uint                        `json:"id"`
	Character  HonkaiCharacterDTO          `json:"character"`
	Cones      []HonkaiLightConeDTO        `json:"cones"`
	Artifacts  []HonkaiArtifactsPresetDTO  `json:"artifacts"`
	Stats      HonkaiStatsDTO              `json:"stats"`
}

type HonkaiLightConeDTO struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Rarity   int    `json:"rarity"`
	Path     string `json:"path"`
	BaseDEF  int32 `json:"base_def"`
	BaseHP   int32 `json:"base_hp"`
	BaseATK  int32   `json:"base_atk"`
	Passive  string `json:"passive"`
}

type HonkaiStatsDTO struct {
	ID           uint   `json:"id"`
	Body         string `json:"body"`
	Feet         string `json:"feet"`
	PlanarSphere string `json:"planar_sphere"`
	LinkRope     string `json:"link_rope"`
	SubStats     string `json:"sub_stats"`
}

type HonkaiArtifactsPresetDTO struct {
	ID     uint              `json:"id"`
	Relics HonkaiRelicsDTO   `json:"relics"`
	Planar HonkaiPlanarDTO   `json:"planar"`
}

type HonkaiRelicsDTO struct {
	ID             uint   `json:"id"`
	Name           string `json:"name"`
	TwoPieceBonus  string `json:"two_piece_bonus"`
	FourPieceBonus string `json:"four_piece_bonus"`
}

type HonkaiPlanarDTO struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	SetBonus  string `json:"set_bonus"`
}