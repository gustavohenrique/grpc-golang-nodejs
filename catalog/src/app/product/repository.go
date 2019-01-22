package product

import (
	"github.com/gustavohenrique/grpc-golang-nodejs/catalog/src/database"
	"github.com/gustavohenrique/grpc-golang-nodejs/catalog/src/interfaces"
)

type ProductRepository struct {
	db database.Database
}

func NewRepository(db database.Database) interfaces.Repository {
	return &ProductRepository{db: db}
}

func (this *ProductRepository) FetchAll() (interface{}, error) {
	var products []Product
	query := "SELECT * FROM products"
	conn := this.db.GetConnection()
	err := conn.Select(&products, query)
	return products, err
}

func (this *ProductRepository) FindByID(id string) (interface{}, error) {
	return nil, nil
}
