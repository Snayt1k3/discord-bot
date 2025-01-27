package interfaces

type Repository[T any] interface {
	Create(model *T) (*T, error)
	GetByID(id uint) (*T, error)
	Update(model *T) error
	Delete(id uint) error
	GetAll() ([]T, error)
}
