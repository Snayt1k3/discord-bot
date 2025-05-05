package interfaces

type ServiceInterface[character, characterBrief, build any] interface {
	GetCharacters() ([]characterBrief, error)
	GetCharacterByID(id string) (character, error)
	GetCharacterBuild(id string) (build, error)
}
