package customer

import (
	"github.com/gustavohenrique/grpc-golang-nodejs/catalog/src/database"
	"github.com/gustavohenrique/grpc-golang-nodejs/catalog/src/interfaces"
)

type CustomerRepository struct {
	db database.Database
}

func NewRepository(db database.Database) interfaces.Repository {
	return &CustomerRepository{db: db}
}

func (this *CustomerRepository) FetchAll() (interface{}, error) {
	return nil, nil
}

func (this *CustomerRepository) FindByID(id string) (interface{}, error) {
	var customer Customer
	query := "SELECT * FROM customers WHERE id = $1"
	conn := this.db.GetConnection()
	err := conn.Get(&customer, query, id)
	return customer, err
}
