package interfaces

type ServiceInterface[character, characterBrief, build any] interface {
	GetCharacters() ([]characterBrief, error)
	GetCharacterByID(id string) (character, error)
	GetCharacterBuilds(id string) ([]build, error)
}
