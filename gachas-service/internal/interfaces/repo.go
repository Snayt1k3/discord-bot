package interfaces

type RepoInterface[character, build any] interface {
	GetCharacters() ([]character, error)
	GetCharacterByID(id string) (character, error)
	GetCharacterBuilds(id string) ([]build, error)
}
