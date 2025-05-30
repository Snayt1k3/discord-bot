package services

import (
	"gachas-service/internal/adapters/storage/models/honkai"
	"gachas-service/internal/dto"
	"gachas-service/internal/interfaces"
	mappers "gachas-service/internal/utils/mappers/honkai"
)

type HonkaiService struct {
	repository interfaces.RepoInterface[honkai.HonkaiCharacter, honkai.HonkaiBuild]
}

func (s *HonkaiService) GetCharacterByID(id string) (dto.HonkaiCharacterDTO, error) {
	data, err := s.repository.GetCharacterByID(id)
	if err != nil {
		return dto.HonkaiCharacterDTO{}, err
	}
	return mappers.MapCharacterToDTO(data), nil
}

func (s *HonkaiService) GetCharacterBuild(id string) (dto.HonkaiBuildDTO, error) {
	buildData, err := s.repository.GetCharacterBuild(id)
	if err != nil {
		return dto.HonkaiBuildDTO{}, err
	}
	return mappers.MapBuildToDTO(buildData), nil
}

func (s *HonkaiService) GetCharacters() ([]dto.HonkaiCharacterBriefDTO, error) {
	characters, err := s.repository.GetCharacters()
	if err != nil {
		return nil, err
	}
	var briefs []dto.HonkaiCharacterBriefDTO
	for _, c := range characters {
		briefs = append(briefs, mappers.MapCharacterBriefToDTO(c))
	}
	return briefs, nil
}

func NewHonkaiService(repo interfaces.RepoInterface[honkai.HonkaiCharacter, honkai.HonkaiBuild]) *HonkaiService {
	return &HonkaiService{
		repository: repo,
	}
}