package repository

import (
	"cabbie/models"
	"errors"
)

type OrderRepository struct {
	MapDatastore map[string]models.Order
}

func (r OrderRepository) CreateNewOrder(order models.Order) error {
	if _, ok := r.MapDatastore[order.OrderID]; ok != false {
		return errors.New("order already exists")
	}
	r.MapDatastore[order.OrderID] = order
	return nil
}

func (r OrderRepository) UpdateOrder(order models.Order) {
	r.MapDatastore[order.OrderID] = order
}

func (r OrderRepository) GetOrderByOrderID(phone string) (models.Order, error) {
	order, ok := r.MapDatastore[phone]
	if ok == false {
		return models.Order{}, errors.New("order not found")
	}
	return order, nil
}

func (r OrderRepository) GetAllOrders() map[string]models.Order {
	return r.MapDatastore
}
