package interfaces

type ServiceInterface[model, build any] interface {
	GetCharacters() ([]model, error)
	GetCharacterByID(id string) (model, error)
	GetCharacterBuild(id string) (build, error)
}
