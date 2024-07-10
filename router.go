package main

import (
	"cabbie/customer"
	"cabbie/driver"
	"github.com/gorilla/mux"
)

func NewRouter(customerService customer.Service, driverService driver.Service) *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/ping", pingHandler).Methods("GET")
	router.HandleFunc("/customer/signup", customer.SignUpHandler(customerService)).Methods("POST")
	router.HandleFunc("/driver/signup", driver.SignUpHandler(driverService)).Methods("POST")
	return router
}
