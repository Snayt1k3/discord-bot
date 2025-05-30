package honkai

import (
	"gachas-service/internal/adapters/storage/models/honkai"
	"gachas-service/internal/dto"
)

func MapCharacterToDTO(c honkai.HonkaiCharacter) dto.HonkaiCharacterDTO {
	return dto.HonkaiCharacterDTO{
		ID:        c.ID,
		Name:      c.Name,
		Rarity:    c.Rarity,
		Path:      c.Path,
		Attribute: c.Attribute,
		Ascension: dto.HonkaiAscensionDTO{
			BossMaterial: c.Ascension.BossMaterial,
			Resource:     mapResourceToDTO(c.Ascension.Resource),
		},
		Traces: dto.HonkaiTracesDTO{
			BossMaterial:    c.Traces.BossMaterial,
			DungeonResource: mapResourceToDTO(c.Traces.DungeonResource),
			MobResource:     mapResourceToDTO(c.Traces.MobResource),
		},
	}
}

func MapCharacterBriefToDTO(c honkai.HonkaiCharacter) dto.HonkaiCharacterBriefDTO {
	return dto.HonkaiCharacterBriefDTO{
		ID:     c.ID,
		Name:   c.Name,
		Rarity: c.Rarity,
	}
}

func MapBuildToDTO(b honkai.HonkaiBuild) dto.HonkaiBuildDTO {
	var cones []dto.HonkaiLightConeDTO
	for _, cone := range b.Cones {
		cones = append(cones, dto.HonkaiLightConeDTO{
			Name:    cone.Name,
			Rarity:  cone.Rarity,
			Path:    cone.Path,
			BaseDEF: cone.BaseDEF,
			BaseHP:  cone.BaseHP,
			BaseATK: cone.BaseATK,
			Passive: cone.Passive,
		})
	}

	var artifacts []dto.HonkaiArtifactsPresetDTO
	for _, a := range b.Artifacts {
		artifacts = append(artifacts, dto.HonkaiArtifactsPresetDTO{
			Relics: dto.HonkaiRelicsDTO{
				Name:           a.Relics.Name,
				TwoPieceBonus:  a.Relics.TwoPieceBonus,
				FourPieceBonus: a.Relics.FourPieceBonus,
			},
			Planar: dto.HonkaiPlanarDTO{
				Name:     a.Planars.Name,
				SetBonus: a.Planars.SetBonus,
			},
		})
	}

	return dto.HonkaiBuildDTO{
		Character: MapCharacterToDTO(b.Character),
		Cones:     cones,
		Artifacts: artifacts,
		Stats: dto.HonkaiStatsDTO{
			Body:         b.Stats.Body,
			Feet:         b.Stats.Feet,
			PlanarSphere: b.Stats.PlanarSphere,
			LinkRope:     b.Stats.LinkRope,
			SubStats:     b.Stats.SubStats,
		},
	}
}

func mapResourceToDTO(r honkai.HonkaiResource) dto.HonkaiResourceDTO {
	return dto.HonkaiResourceDTO{
		Uncommon: r.Uncommon,
		Rare:     r.Rare,
		Epic:     r.Epic,
	}
}
