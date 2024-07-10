package repository

import (
	"cabbie/models"
	"errors"
)

type DriverRepository struct {
	MapDatastore map[string]models.Driver
}

func (r *DriverRepository) AddDriver(driver models.Driver) error {
	if _, ok := r.MapDatastore[driver.Phone]; ok != false {
		return errors.New("driver already exists")
	}
	r.MapDatastore[driver.Phone] = driver
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