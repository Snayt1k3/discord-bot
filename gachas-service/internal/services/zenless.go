package services

import (
	"gachas-service/internal/adapters/storage/models/zenless"
	"gachas-service/internal/dto"
	"gachas-service/internal/interfaces"
	mappers "gachas-service/internal/utils/mappers/zenless"
)

type ZenlessService struct {
	repository interfaces.RepoInterface[zenless.ZenlessCharacter, zenless.ZenlessBuild]
}

func (s *ZenlessService) GetCharacterByID(id string) (dto.ZenlessCharacterFull, error) {
	data, err := s.repository.GetCharacterByID(id)
	if err != nil {
		return dto.ZenlessCharacterFull{}, err
	}

	dto := mappers.MapCharacterToDTO(data)
	return dto, nil
}

func (s *ZenlessService) GetCharacterBuild(id string) (dto.ZenlessBuildDTO, error) {
	buildData, err := s.repository.GetCharacterBuild(id)

	if err != nil {
		return dto.ZenlessBuildDTO{}, err
	}

	dto := mappers.MapBuildToDTO(&buildData)
	return dto, nil
}

func (s *ZenlessService) GetCharacters() ([]dto.ZenlessCharacterDTO, error) {
	characters, err := s.repository.GetCharacters()

	if err != nil {
		return nil, err
	}

	var result []dto.ZenlessCharacterDTO

	for _, c := range characters {
		result = append(result, mappers.MapCharacterToShortDTO(c))
	}
	return result, nil
}

func NewZenlessService(repo interfaces.RepoInterface[zenless.ZenlessCharacter, zenless.ZenlessBuild]) *ZenlessService {
	return &ZenlessService{
		repository: repo,
	}
}
