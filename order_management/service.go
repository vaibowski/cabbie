package order_management

import (
	"cabbie/driver"
	"cabbie/models"
)

type Repository interface {
	CreateNewOrder(order models.Order) error
}

type DriverService interface {
	AllocateDriver(pickup models.Location, serviceType models.ServiceTypeEnum)
}

type Service struct {
	OrderRepository   Repository
	AllocationService driver.AllocationService
}

func (svc Service) CreateNewOrder(order models.Order) (models.Order, error) {
	return models.Order{}, nil
}

func NewService(orderRepository Repository, allocationService driver.AllocationService) Service {
	return Service{
		OrderRepository:   orderRepository,
		AllocationService: allocationService,
	}
}
