package services


import (
	"gachas-service/internal/dto"
	"gachas-service/internal/interfaces"
)


type GenshinService struct {
	repository interfaces.RepoInterface[dto.GenshinCharacter, dto.GenshinBuild]
}

func (s *GenshinService) GetCharacter(id string) (*dto.GenshinCharacter, error) {
	return nil, nil
}

func (s *GenshinService) GetBuild(id string) (*dto.GenshinBuild, error) {
	return nil, nil
}

func (s *GenshinService) GetCharacters() ([]dto.GenshinCharacter, error) {
	return nil, nil
}



func NewGenshinService(repo interfaces.RepoInterface[dto.GenshinCharacter, dto.GenshinBuild]) *GenshinService {
	return &GenshinService{
		repository: repo,
	}
}