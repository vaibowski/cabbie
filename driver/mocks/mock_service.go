package mocks

import (
	"cabbie/models"
	"github.com/stretchr/testify/mock"
)

type MockService struct {
	mock.Mock
}

func (service *MockService) FetchDriver(driverID string) (models.Driver, error) {
	args := service.Called(driverID)
	return args.Get(0).(models.Driver), args.Error(1)
}

func (service *MockService) UpdateDriver(driver models.Driver) {
	return
}

func (service *MockService) CreateNewDriver(customer models.Driver) (string, error) {
	args := service.Called(customer)
	return args.String(0), args.Error(1)
}
