package product_test

import (
	"testing"
	"errors"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/gustavohenrique/grpc-golang-nodejs/catalog/src/interfaces"
	"github.com/gustavohenrique/grpc-golang-nodejs/catalog/src/app/product"
	"github.com/gustavohenrique/grpc-golang-nodejs/catalog/src/app/customer"
)

type CustomerRepositoryMock struct {
	mock.Mock
}

func (m *CustomerRepositoryMock) FetchAll() (interface{}, error) {
	return nil, nil
}

func (m *CustomerRepositoryMock) FindByID(id string) (interface{}, error) {
	args := m.Called(id)
	c := customer.Customer{}
	return c, args.Error(1)
}

func TestShoulFetchAllProductsWhenCustomerIdIsEmpty(t *testing.T) {
    customerRepositoryMock := new(CustomerRepositoryMock)
	repositories := map[string]interfaces.Repository{
		"product": nil,
		"customer": customerRepositoryMock,
	}
	service := product.NewService(repositories)
	_, err := service.FindCustomerByID("")
	assert.NotNil(t, err)

	customerRepositoryMock.AssertNotCalled(t, "FindByID")
}

func TestShoulFetchAllProductsWhenCustomerIdIsNotFound(t *testing.T) {
	customerRepositoryMock := new(CustomerRepositoryMock)
	customerRepositoryMock.On("FindByID", mock.Anything).Return(customer.Customer{}, errors.New("Not found"))
	repositories := map[string]interfaces.Repository{
		"product": nil,
		"customer": customerRepositoryMock,
	}
	service := product.NewService(repositories)
	_, err := service.FindCustomerByID("9999999")
	assert.NotNil(t, err)

	customerRepositoryMock.AssertCalled(t, "FindByID", "9999999")
}