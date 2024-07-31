package models

import "time"

type Driver struct {
	DriverID     string          `json:"driverID"`
	Name         string          `json:"name"`
	Phone        string          `json:"phone"`
	Email        string          `json:"email"`
	Password     string          `json:"password"`
	LastLocation Location        `json:"lastLocation"`
	ServiceType  ServiceTypeEnum `json:"serviceType"`
	CreatedAt    time.Time       `json:"createdAt"`
}
