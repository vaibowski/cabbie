package mocks

import (
	"cabbie/models"
	"github.com/stretchr/testify/mock"
)

type MockService struct {
	mock.Mock
}

func (service *MockService) CreateNewCustomer(customer models.Customer) (string, error) {
	args := service.Called(customer)
	return args.String(0), args.Error(1)
}
