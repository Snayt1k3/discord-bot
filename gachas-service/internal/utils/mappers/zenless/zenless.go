package zenless

import (
	"gachas-service/internal/adapters/storage/models/zenless"
	"gachas-service/internal/dto"
)

// Character short DTO (для списка)
func MapCharacterToShortDTO(c zenless.ZenlessCharacter) dto.ZenlessCharacterDTO {
	return dto.ZenlessCharacterDTO{
		ID:        c.ID,
		Name:      c.Name,
		Rank:      c.Rank,
		Attribute: c.Attribute,
		Faction:   c.Faction,
	}
}

// Полный Character с Ascension и Nodes
func MapCharacterToDTO(c zenless.ZenlessCharacter) dto.ZenlessCharacterFull {
	return dto.ZenlessCharacterFull{
		ID:        c.ID,
		Name:      c.Name,
		Specialty: c.Specialty,
		Rank:      c.Rank,
		Attribute: c.Attribute,
		Faction:   c.Faction,
		Ascension: dto.ZenlessAscensionDTO{
			Common: c.Ascension.Common,
			Rare:   c.Ascension.Rare,
			Epic:   c.Ascension.Epic,
		},
		Nodes: dto.ZenlessNodesDTO{
			HuntBossMaterial:   c.Nodes.HuntBossMaterial,
			ExpertBossMaterial: c.Nodes.ExpertBossMaterial,
			Resource: dto.ZenlessAscensionDTO{
				Common: c.Nodes.Resource.Common,
				Rare:   c.Nodes.Resource.Rare,
				Epic:   c.Nodes.Resource.Epic,
			},
		},
	}
}

func MapBuildToDTO(b *zenless.ZenlessBuild) dto.ZenlessBuildDTO {
	var weapons []dto.ZenlessWeaponDTO
	for _, w := range b.Weapons {
		weapons = append(weapons, dto.ZenlessWeaponDTO{
			Name:     w.Name,
			BaseATK:  w.BaseATK,
			Rarity:   w.Rarity,
			Type:     w.Type,
			Passive:  w.Passive,
			SubStat:  w.SubStat,
			SubValue: w.SubValue,
		})
	}

	var discs []dto.ZenlessDiscDTO
	for _, preset := range b.Discs {
		discs = append(discs, dto.ZenlessDiscDTO{
			TwoPiece: dto.ZenlessDiscSetDTO{
				Name:           preset.TwoPieceSet.Name,
				TwoPieceBonus:  preset.TwoPieceSet.TwoPieceBonus,
				FourPieceBonus: preset.TwoPieceSet.FourPieceBonus,
			},
			FourPiece: dto.ZenlessDiscSetDTO{
				Name:           preset.FourPieceSet.Name,
				TwoPieceBonus:  preset.FourPieceSet.TwoPieceBonus,
				FourPieceBonus: preset.FourPieceSet.FourPieceBonus,
			},
		})
	}

	return dto.ZenlessBuildDTO{
		Character: MapCharacterToDTO(b.Character),
		Weapons:   weapons,
		Discs:     discs,
		Stats: dto.ZenlessStatsDTO{
			FourthDisc:       b.Stats.FourthDisc,
			FifthDisc:        b.Stats.FifthDisc,
			SixthDisc:        b.Stats.SixthDisc,
			SubStatsPriority: b.Stats.SubStatsPriority,
		},
	}
}
