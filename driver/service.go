package driver

import (
	"cabbie/models"
	"errors"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"log"
	"time"
)

type Repository interface {
	AddDriver(driver models.Driver) error
	GetDriverByPhone(phone string) (models.Driver, error)
	GetDriverByDriverID(driverID string) (models.Driver, error)
	UpdateDriver(driver models.Driver)
}

type allocator interface {
	AllocateDriver(pickup models.Location, serviceType models.ServiceTypeEnum) (string, error)
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

func (service Service) FetchDriver(driverID string) (models.Driver, error) {
	driver, err := service.DriverRepository.GetDriverByDriverID(driverID)
	if err != nil {
		log.Printf("driver validation failed for driverID: %s with %s", driverID, err.Error())
		return models.Driver{}, errors.New("driver validation failed for driverID: " + driverID)
	} else {
		return driver, nil
	}
}

func (service Service) UpdateDriver(driver models.Driver) {
	service.DriverRepository.UpdateDriver(driver)
}

func NewService(driverRepository Repository) Service {
	return Service{DriverRepository: driverRepository}
}
