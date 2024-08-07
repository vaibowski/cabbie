package repository

import (
	"cabbie/models"
	"errors"
	"log"
)

type DriverRepository struct {
	MapDatastore map[string]models.Driver
}

func (r *DriverRepository) AddDriver(driver models.Driver) error {
	if _, ok := r.MapDatastore[driver.Phone]; ok != false {
		return errors.New("driver already exists")
	}
	r.MapDatastore[driver.Phone] = driver
	r.MapDatastore[driver.DriverID] = driver
	log.Printf("All drivers in datastore: %v", r.MapDatastore)
	return nil
}

// GetDriverByPhone to be used later on for login
func (r *DriverRepository) GetDriverByPhone(phone string) (models.Driver, error) {
	driver, ok := r.MapDatastore[phone]
	if ok == false {
		return models.Driver{}, errors.New("driver not found")
	}
	return driver, nil
}

func (r *DriverRepository) GetDriverByDriverID(driverID string) (models.Driver, error) {
	driver, ok := r.MapDatastore[driverID]
	if ok == false {
		return models.Driver{}, errors.New("driver not found")
	}
	return driver, nil
}

func (r *DriverRepository) UpdateDriver(driver models.Driver) {
	r.MapDatastore[driver.DriverID] = driver
	r.MapDatastore[driver.Phone] = driver
}
