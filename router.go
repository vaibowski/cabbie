package main

import (
	"cabbie/customer"
	"cabbie/driver"
	"cabbie/estimate"
	"github.com/gorilla/mux"
)

func NewRouter(customerService customer.Service, driverService driver.Service, estimateService estimate.Service) *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/ping", pingHandler).Methods("GET")
	router.HandleFunc("/customer/signup", customer.SignUpHandler(customerService)).Methods("POST")
	router.HandleFunc("/driver/signup", driver.SignUpHandler(driverService)).Methods("POST")
	router.HandleFunc("/transport/estimate", estimate.Handler(estimateService))
	return router
}
