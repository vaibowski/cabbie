package customer

import (
	"cabbie/models"
	"github.com/google/uuid"
	"log"
	"time"
)

type Repository interface {
	AddCustomer(customer models.Customer) error
	GetCustomerByID(customerID string) (models.Customer, error)
}

type Service struct {
	CustomerRepository Repository
}

func (service Service) CreateNewCustomer(customer models.Customer) (string, error) {
	customerID := uuid.New().String()
	customer.CustomerID = customerID
	customer.CreatedAt = time.Now()
	err := service.CustomerRepository.AddCustomer(customer)
	if err != nil {
		return "", err
	}
	log.Printf("customer successfully created with id %s", customerID)
	return customerID, nil
}

func NewService(customerRepository Repository) Service {
	return Service{CustomerRepository: customerRepository}
}
