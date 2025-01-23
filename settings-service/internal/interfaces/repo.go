package interfaces

type Repository[T any] interface {
    Create(user T) (int, error)      
    GetByID(id int) (T, error)      
    Update(user T) error            
    Delete(id int) error               
    GetAll() ([]T, error)           
}
