package services

import (
	"gachas-service/internal/adapters/storage/models/wuwa"
	"gachas-service/internal/dto"
	"gachas-service/internal/interfaces"
	mappers "gachas-service/internal/utils/mappers/wuwa"
)

type WuwaService struct {
	repository interfaces.RepoInterface[wuwa.Character, wuwa.Build]
}

func (s *WuwaService) GetCharacterByID(id string) (dto.WuwaCharacterFull, error) {
	data, err := s.repository.GetCharacterByID(id)

	if err != nil {
		return dto.WuwaCharacterFull{}, err
	}

	dto := mappers.MapCharacterToDTO(data)

	return dto, nil

}

func (s *WuwaService) GetCharacterBuild(id string) (dto.WuwaCharacterBuild, error) {
	var build dto.WuwaCharacterBuild

	buildData, err := s.repository.GetCharacterBuild(id)

	if err != nil {
		return build, err
	}

	build = mappers.MapBuildToDTO(buildData)

	return build, nil
}

func (s *WuwaService) GetCharacters() ([]dto.WuwaCharacterShort, error) {

	characters, err := s.repository.GetCharacters()

	if err != nil {
		return nil, err
	}
	var characterBriefs []dto.WuwaCharacterShort

	for _, character := range characters {
		characterBriefs = append(characterBriefs, mappers.MapCharacterToShortDTO(character))
	}
	return characterBriefs, nil
}

func NewWuwaService(repo interfaces.RepoInterface[wuwa.Character, wuwa.Build]) *WuwaService{
	return &WuwaService{
		repository: repo,
	}
}
