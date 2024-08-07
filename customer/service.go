package customer

import (
	"cabbie/models"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"log"
	"time"
)

type Repository interface {
	AddCustomer(customer models.Customer) error
	GetCustomerByPhone(customerID string) (models.Customer, error)
}

type Service struct {
	CustomerRepository Repository
}

func (service Service) CreateNewCustomer(customer models.Customer) (string, error) {
	customerID := uuid.New().String()
	customer.CustomerID = customerID
	customer.CreatedAt = time.Now()
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(customer.Password), 10)
	if err != nil {
		return "", err
	}
	customer.Password = string(passwordHash)
	err = service.CustomerRepository.AddCustomer(customer)
	if err != nil {
		return "", err
	}
	log.Printf("customer successfully created with id %s", customerID)
	return customerID, nil
}

func NewService(customerRepository Repository) Service {
	return Service{CustomerRepository: customerRepository}
}
