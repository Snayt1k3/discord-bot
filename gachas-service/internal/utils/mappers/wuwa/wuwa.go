package wuwa

import (
	models "gachas-service/internal/adapters/storage/models/wuwa"
	"gachas-service/internal/dto"
)

// --- Основные мапперы ---

func MapCharacterToDTO(c models.WuwaCharacter) dto.WuwaCharacterFull {
	return dto.WuwaCharacterFull{
		ID:      c.ID,
		Name:    c.Name,
		Element: c.Element,
		WeaponType: c.WeaponType,
		Rarity:  c.Rarity,
		Ascension: dto.WuwaAscension{
			LocalSpecialty: c.Ascension.LocalSpecialty,
			BossMaterial:   c.Ascension.BossMaterial,
			MobMaterial:    mapResource(c.Ascension.MobMaterial),
		},
		Talents: dto.WuwaTalent{
			DungeonMaterial: mapResource(c.Talents.DungeonMaterial),
			MobMaterial:     mapResource(c.Talents.MobMaterial),
			BossMaterial: c.Talents.BossMaterial,
		},
	}
}

func MapCharacterToShortDTO(c models.WuwaCharacter) dto.WuwaCharacterShort {
	return dto.WuwaCharacterShort{
		ID:      c.ID,
		Name:    c.Name,
		Element: c.Element,
		Rarity:  c.Rarity,
	}
}

func MapBuildToDTO(b models.WuwaBuild) dto.WuwaCharacterBuild {
	return dto.WuwaCharacterBuild{
		Character:       MapCharacterToDTO(b.Character),
		Weapons:         mapWeapons(b.Weapons),
		Echoes:          mapEchoes(b.Echoes),
		Stats:           mapStats(b.Stats),
		BestPrimaryEcho: b.BestPrimaryEcho,
	}
}

// --- Вспомогательные функции ---

func mapWeapon(w models.WuwaWeapon) dto.WuwaWeapon {
	return dto.WuwaWeapon{
		Name:       w.Name,
		WeaponType: w.WeaponType,
		Rarity:     w.Rarity,
		Passive:    w.Passive,
		BaseATK:    w.BaseATK,
		SubStat:    w.SubStat,
		SubValue:   w.SubValue,
	}
}

func mapWeapons(ws []models.WuwaWeapon) []dto.WuwaWeapon {
	result := make([]dto.WuwaWeapon, len(ws))
	for i, w := range ws {
		result[i] = mapWeapon(w)
	}
	return result
}

func mapResource(r models.WuwaResource) dto.WuwaResource {
	return dto.WuwaResource{
		UncommonName:  r.UncommonName,
		RareName:      r.RareName,
		EpicName:      r.EpicName,
		LegendaryName: r.LegendaryName,
	}
}

func mapEchoes(es []models.WuwaEchoes) []dto.WuwaEcho {
	result := make([]dto.WuwaEcho, len(es))
	for i, e := range es {
		result[i] = dto.WuwaEcho{
			Name:          e.Name,
			TwoPieceBonus: e.TwoPieceBonus,
			FullSetBonus:  e.FullSetBonus,
		}
	}
	return result
}

func mapStats(s models.WuwaStats) dto.WuwaStats {
	return dto.WuwaStats{
		FourCostEchoStat:  s.FourCostEchoStat,
		ThreeCostEchoStat: s.ThreeCostEchoStat,
		SubStatsPriority:  s.SubStatsPriority,
		OneCostEchoStat:   s.OneCostEchoStat,
	}
}
