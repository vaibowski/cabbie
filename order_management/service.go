package order_management

import (
	"cabbie/models"
	"github.com/google/uuid"
	"log"
	"time"
)

type Repository interface {
	CreateNewOrder(order models.Order) error
	GetOrderByOrderID(orderID string) (models.Order, error)
	GetAllOrders() map[string]models.Order
	UpdateOrder(order models.Order)
}

type allocationService interface {
	AllocateDriver(origin models.Location, serviceType models.ServiceTypeEnum) (string, error)
	UnsetLocation(driverID string, serviceType models.ServiceTypeEnum, lastLocation models.Location) error
}

type Service struct {
	OrderRepository   Repository
	AllocationService allocationService
}

func (svc Service) CreateNewOrder(order models.Order) (models.Order, error) {
	orderID := uuid.New().String()
	order.OrderID = orderID
	order.CreatedAt = time.Now()
	order.Status = models.CREATED
	err := svc.OrderRepository.CreateNewOrder(order)
	if err != nil {
		log.Printf("error creating order: %v", err)
		return models.Order{}, err
	}
	log.Printf("created new order: %v", order)

	time.Sleep(20)

	// order has been created, now we will assign a driver
	driverID, err := svc.AllocationService.AllocateDriver(order.Origin, order.ServiceType)
	if err != nil {
		log.Printf("error allocating driver for orderID: %s, with err: %s", orderID, err.Error())
		return order, err
	}
	order.DriverID = driverID
	order.Status = models.DRIVER_ASSIGNED
	svc.OrderRepository.UpdateOrder(order)
	return order, nil
}

func (svc Service) StartOrder(orderID string) (models.Order, error) {
	order, err := svc.OrderRepository.GetOrderByOrderID(orderID)
	if err != nil {
		log.Printf("order not found: %v", err)
		return models.Order{}, err
	}
	order.Status = models.RIDE_IN_PROGRESS
	order.PickupTime = time.Now()
	svc.OrderRepository.UpdateOrder(order)
	return order, nil
}

func (svc Service) CompleteOrder(orderID string) (models.Order, error) {
	order, err := svc.OrderRepository.GetOrderByOrderID(orderID)
	if err != nil {
		log.Printf("order not found: %v", err)
		return models.Order{}, err
	}
	order.Status = models.COMPLETED
	order.DropOffTime = time.Now()
	svc.OrderRepository.UpdateOrder(order)
	return order, nil
}

func (svc Service) FetchOrder(orderID string) (models.Order, error) {
	return svc.OrderRepository.GetOrderByOrderID(orderID)
}

func (svc Service) FetchAllOrders() map[string]models.Order {
	return svc.OrderRepository.GetAllOrders()
}

func NewService(orderRepository Repository, allocationService allocationService) Service {
	return Service{
		OrderRepository:   orderRepository,
		AllocationService: allocationService,
	}
}
