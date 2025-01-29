package interfaces

type Repository[T any] interface {
	Create(model *T) (*T, error)
	Updates(id string, fields map[string]interface{}) error
	Delete(id uint) error
	GetAll() ([]T, error)
	Filter(filters map[string]interface{}) ([]T, error)
}
