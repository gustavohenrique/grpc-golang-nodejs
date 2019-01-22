package interfaces

type Repository interface {
	FetchAll() (interface{}, error)
	FindByID(id string) (interface{}, error)
}