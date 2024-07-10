package driver

import (
	"cabbie/models"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"log"
	"time"
)

type Repository interface {
	AddDriver(driver models.Driver) error
	GetDriverByPhone(driverID string) (models.Driver, error)
}

type Service struct {
	DriverRepository Repository
}

func (service Service) CreateNewDriver(driver models.Driver) (string, error) {
	driverID := uuid.New().String()
	driver.DriverID = driverID
	driver.CreatedAt = time.Now()
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(driver.Password), 10)
	if err != nil {
		return "", err
	}
	driver.Password = string(passwordHash)
	err = service.DriverRepository.AddDriver(driver)
	if err != nil {
		return "", err
	}
	log.Printf("driver successfully created with id %s", driverID)
	return driverID, nil
}

func NewService(driverRepository Repository) Service {
	return Service{DriverRepository: driverRepository}
}
