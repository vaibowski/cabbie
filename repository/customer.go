package repository

import (
	"cabbie/models"
	"errors"
	"log"
)

type CustomerRepository struct {
	MapDatastore map[string]models.Customer
}

func (r *CustomerRepository) AddCustomer(customer models.Customer) error {
	if _, ok := r.MapDatastore[customer.Phone]; ok != false {
		return errors.New("customer already exists")
	}
	r.MapDatastore[customer.Phone] = customer
	r.MapDatastore[customer.CustomerID] = customer
	log.Printf("All customers in datastore: %v", r.MapDatastore)
	return nil
}

// GetCustomerByPhone to be used later on for login
func (r *CustomerRepository) GetCustomerByPhone(phone string) (models.Customer, error) {
	customer, ok := r.MapDatastore[phone]
	if ok == false {
		return models.Customer{}, errors.New("customer not found")
	}
	return customer, nil
}

func (r *CustomerRepository) GetCustomerByCustomerID(customerID string) (models.Customer, error) {
	customer, ok := r.MapDatastore[customerID]
	if ok == false {
		return models.Customer{}, errors.New("customer not found")
	}
	return customer, nil
}
