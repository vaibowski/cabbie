package repository

import (
	"cabbie/models"
)

type CustomerRepository struct {
	Datastore map[string]models.Customer
}

func (r CustomerRepository) AddCustomer(customer models.Customer) error {
	// TODO: add repo implementation
	return nil
}

func (r CustomerRepository) GetCustomerByID(customerID string) (models.Customer, error) {
	// TODO: add repo implementation
	return models.Customer{}, nil
}
