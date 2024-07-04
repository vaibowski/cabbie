package models

import "time"

type Customer struct {
	CustomerID string    `json:"customerID"`
	Name       string    `json:"name"`
	Email      string    `json:"email"`
	Password   string    `json:"password"`
	Phone      string    `json:"phone"`
	CreatedAt  time.Time `json:"createdAt"`
}
