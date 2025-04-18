package services

import (
	"gachas-service/internal/adapters/storage/models/genshin"
	"gachas-service/internal/dto"
	"gachas-service/internal/interfaces"
	"gachas-service/internal/utils/mappers"
)

type GenshinService struct {
	repository interfaces.RepoInterface[genshin.Character, genshin.Build]
}

func (s *GenshinService) GetCharacterByID(id string) (*dto.GenshinCharacter, error) {
	data, err := s.repository.GetCharacterByID(id)

	if err != nil {
		return nil, err
	}

	dto := mappers.MapCharacterToDTO(&data)

	return &dto, nil

}

func (s *GenshinService) GetCharacterBuild(id string) (*dto.GenshinBuild, error) {

	buildData, err := s.repository.GetCharacterBuild(id)

	if err != nil {
		return nil, err
	}

	buildDto := mappers.MapBuildToDTO(&buildData)

	return buildDto, nil
}

func (s *GenshinService) GetCharacters() (*[]dto.GenshinCharacterBrief, error) {

	characters, err := s.repository.GetCharacters()

	if err != nil {
		return nil, err
	}
	var characterBriefs []dto.GenshinCharacterBrief

	for _, character := range characters {
		characterBriefs = append(characterBriefs, mappers.MapCharacterBriefToDTO(&character))
	}
	return &characterBriefs, nil
}

func NewGenshinService(repo interfaces.RepoInterface[genshin.Character, genshin.Build]) *GenshinService {
	return &GenshinService{
		repository: repo,
	}
}
