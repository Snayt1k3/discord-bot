package genshin

import (
	"gachas-service/internal/adapters/storage/models/genshin"
	"gachas-service/internal/dto"
)

func MapBuildToDTO(build genshin.GenshinBuild) dto.GenshinBuild {
	return dto.GenshinBuild{
		ID:        build.ID,
		Character: MapCharacterToDTO(build.Character),
		Teams:     MapTeamsToDTO(build.Teams),
		Artifacts: MapArtifactsToDTO(build.Artifacts),
		Stats:     MapStatsToDTO(build.Stats),
		Weapons:   MapWeaponsToDTO(build.Weapons),
	}
}

func MapWeaponsToDTO(weapons []genshin.GenshinWeapon) []dto.GenshinWeapon {
	weaponsDto := make([]dto.GenshinWeapon, len(weapons))
	for i, weapon := range weapons {
		weaponsDto[i] = dto.GenshinWeapon{
			ID:       weapon.ID,
			Name:     weapon.Name,
			Type:     weapon.Type,
			Rarity:   weapon.Rarity,
			SubStat:  weapon.SubStat,
			BaseATK:  weapon.BaseATK,
			SubValue: weapon.SubValue,
			Passive:  weapon.Passive,
		}
	}
	return weaponsDto
}

func MapStatsToDTO(stats genshin.GenshinStats) dto.GenshinStats {
	return dto.GenshinStats{
		ID:               stats.ID,
		Sands:            stats.Sands,
		Goblet:           stats.Goblet,
		Circlet:          stats.Circlet,
		SubStatsPriority: stats.SubStatsPriority,
	}
}

func MapArtifactsToDTO(artifacts []genshin.GenshinArtifact) []dto.GenshinArtifact {
	artifactsDTO := make([]dto.GenshinArtifact, len(artifacts))
	for i, artifact := range artifacts {
		artifactsDTO[i] = dto.GenshinArtifact{
			ID:             artifact.ID,
			Name:           artifact.Name,
			TwoPieceBonus:  artifact.TwoPieceBonus,
			FourPieceBonus: artifact.FourPieceBonus,
		}
	}
	return artifactsDTO
}

func MapTeamsToDTO(teams []genshin.GenshinTeam) []dto.GenshinTeam {
	teamsDTO := make([]dto.GenshinTeam, len(teams))
	for i, team := range teams {
		teamsDTO[i] = dto.GenshinTeam{
			ID:         team.ID,
			Characters: make([]dto.GenshinCharacter, len(team.Characters)),
		}
		for j, character := range team.Characters {
			teamsDTO[i].Characters[j] = MapCharacterToDTO(character)
		}
	}
	return teamsDTO
}

func MapCharacterToDTO(character genshin.GenshinCharacter) dto.GenshinCharacter {
	return dto.GenshinCharacter{
		ID:         character.ID,
		Name:       character.Name,
		WeaponType: character.WeaponType,
		Region:     character.Region,
		Rarity:     character.Rarity,
		BaseStat:   character.BaseStat,
		Ascension: dto.GenshinAscensionMaterials{
			ID:             character.Ascension.ID,
			Gem:            character.Ascension.Gem,
			LocalSpecialty: character.Ascension.LocalSpecialty,
			BossDrops:      character.Ascension.BossDrops,
		},
		Talents: dto.GenshinTalentMaterials{
			ID:        character.Talents.ID,
			BossDrops: character.Talents.BossDrops,
			Books: dto.GenshinBooks{
				ID:       character.Talents.Books.ID,
				Common:   character.Talents.Books.Common,
				Uncommon: character.Talents.Books.Uncommon,
				Rare:     character.Talents.Books.Rare,
				Weekdays: character.Talents.Books.Weekdays,
			},

			TalentPriority: character.Talents.TalentPriority,
		},
		CommonMaterials: dto.GenshinCommonMaterials{
			ID:       character.CommonMaterials.ID,
			Common:   character.CommonMaterials.Common,
			Uncommon: character.CommonMaterials.Uncommon,
			Rare:     character.CommonMaterials.Rare,
		},
		Element: character.Element,
	}
}

func MapCharacterBriefToDTO(character genshin.GenshinCharacter) dto.GenshinCharacterBrief {
	return dto.GenshinCharacterBrief{
		ID:         character.ID,
		Name:       character.Name,
		WeaponType: character.WeaponType,
		Region:     character.Region,
		Rarity:     character.Rarity,
		Element:    character.Element,
	}
}
