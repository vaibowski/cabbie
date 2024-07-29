package repository

import (
	"cabbie/models"
	"errors"
)

type OrderRepository struct {
	MapDatastore map[string]models.Order
}

func (r *OrderRepository) CreateNewOrder(order models.Order) error {
	if _, ok := r.MapDatastore[order.OrderID]; ok != false {
		return errors.New("order already exists")
	}
	r.MapDatastore[order.OrderID] = order
	return nil
}

// GetOrderByOrderID GetOrderByPhone to be used later on for login
func (r *OrderRepository) GetOrderByOrderID(phone string) (models.Order, error) {
	order, ok := r.MapDatastore[phone]
	if ok == false {
		return models.Order{}, errors.New("order not found")
	}
	return order, nil
}
