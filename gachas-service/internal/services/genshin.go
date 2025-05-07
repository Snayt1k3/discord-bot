package services

import (
	"gachas-service/internal/adapters/storage/models/genshin"
	"gachas-service/internal/dto"
	"gachas-service/internal/interfaces"
	mappers "gachas-service/internal/utils/mappers/genshin"
)

type GenshinService struct {
	repository interfaces.RepoInterface[genshin.Character, genshin.Build]
}

func (s *GenshinService) GetCharacterByID(id string) (dto.GenshinCharacter, error) {
	data, err := s.repository.GetCharacterByID(id)

	if err != nil {
		return dto.GenshinCharacter{}, err
	}

	dto := mappers.MapCharacterToDTO(data)

	return dto, nil

}

func (s *GenshinService) GetCharacterBuild(id string) (dto.GenshinBuild, error) {
	var build dto.GenshinBuild

	buildData, err := s.repository.GetCharacterBuild(id)

	if err != nil {
		return build, err
	}

	build = mappers.MapBuildToDTO(buildData)

	return build, nil
}

func (s *GenshinService) GetCharacters() ([]dto.GenshinCharacterBrief, error) {

	characters, err := s.repository.GetCharacters()

	if err != nil {
		return nil, err
	}
	var characterBriefs []dto.GenshinCharacterBrief

	for _, character := range characters {
		characterBriefs = append(characterBriefs, mappers.MapCharacterBriefToDTO(character))
	}
	return characterBriefs, nil
}

func NewGenshinService(repo interfaces.RepoInterface[genshin.Character, genshin.Build]) *GenshinService {
	return &GenshinService{
		repository: repo,
	}
}
