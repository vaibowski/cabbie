package repository

import (
	"cabbie/models"
	"errors"
)

type CustomerRepository struct {
	MapDatastore map[string]models.Customer
}

func (r *CustomerRepository) AddCustomer(customer models.Customer) error {
	if _, ok := r.MapDatastore[customer.Phone]; ok != false {
		return errors.New("customer already exists")
	}
	r.MapDatastore[customer.Phone] = customer
	return nil
}

func (r *CustomerRepository) GetCustomerByPhone(phone string) (models.Customer, error) {
	customer, ok := r.MapDatastore[phone]
	if ok == false {
		return models.Customer{}, errors.New("customer not found")
	}
	return customer, nil
}
