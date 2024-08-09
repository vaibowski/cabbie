package models

import "time"

type Order struct {
	OrderID     string          `json:"orderID"`
	Origin      Location        `json:"origin"`
	Destination Location        `json:"destination"`
	CustomerID  string          `json:"customerID"`
	DriverID    string          `json:"driverID"`
	ServiceType ServiceTypeEnum `json:"serviceType"`
	Status      OrderStatusEnum `json:"status"`
	PickupTime  time.Time       `json:"pickupTime"`
	DropOffTime time.Time       `json:"dropOffTime"`
	TotalPrice  float64         `json:"totalPrice"`
	PaymentID   string          `json:"paymentID"`
	CreatedAt   time.Time       `json:"createdAt"`
}

type ServiceTypeEnum int

const (
	UNKNOWN ServiceTypeEnum = iota
	BIKE
	CAR
	SEDAN
	SUV
)

type OrderStatusEnum int

const (
	UNKNOWN_STATUS OrderStatusEnum = iota
	CREATED
	DRIVER_ASSIGNED
	RIDE_IN_PROGRESS
	COMPLETED
)
