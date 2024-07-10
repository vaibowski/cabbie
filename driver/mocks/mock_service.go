package mocks

import (
	"cabbie/models"
	"github.com/stretchr/testify/mock"
)

type MockService struct {
	mock.Mock
}

func (service *MockService) CreateNewDriver(customer models.Driver) (string, error) {
	args := service.Called(customer)
	return args.String(0), args.Error(1)
}
