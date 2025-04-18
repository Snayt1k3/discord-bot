package interfaces

type RepoInterface[character, build any] interface {
	GetCharacters() ([]character, error)
	GetCharacterByID(id string) (character, error)
	GetCharacterBuild(id string) (build, error)
}
