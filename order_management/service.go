package order_management

import (
	"cabbie/models"
)

type Repository interface {
	CreateNewOrder(order models.Order) error
}

type allocationService interface {
	AllocateDriver(pickup models.Location, serviceType models.ServiceTypeEnum) (string, error)
}

type Service struct {
	OrderRepository   Repository
	AllocationService allocationService
}

func (svc Service) CreateNewOrder(order models.Order) (models.Order, error) {
	return models.Order{}, nil
}

func NewService(orderRepository Repository, allocationService allocationService) Service {
	return Service{
		OrderRepository:   orderRepository,
		AllocationService: allocationService,
	}
}
