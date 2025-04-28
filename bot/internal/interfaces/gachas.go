package interfaces

import (
	dtoGachas "bot/internal/dto/gachas"
)


type GachasAdapter interface {
	GetGenshinCharacters() ([]dtoGachas.GenshinCharacterBrief, error)
	GetGenshinCharacter(id uint) (dtoGachas.GenshinCharacter, error)
	GetGenshinBuild(id uint) (dtoGachas.GenshinBuild, error)

}