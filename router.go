package main

import (
	"cabbie/customer"
	"cabbie/driver"
	"cabbie/estimate"
	"cabbie/order_management"
	"github.com/gorilla/mux"
)

func NewRouter(customerService customer.Service, driverService driver.Service, estimateService estimate.Service, allocationService driver.AllocationService, orderManagementService order_management.Service) *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/ping", pingHandler).Methods("GET")
	router.HandleFunc("/customer/signup", customer.SignUpHandler(customerService)).Methods("POST")
	router.HandleFunc("/driver/signup", driver.SignUpHandler(driverService)).Methods("POST")
	router.HandleFunc("/transport/estimate", estimate.Handler(estimateService)).Methods("GET")
	router.HandleFunc("/transport/create_order", order_management.CreateOrderHandler(orderManagementService)).Methods("POST")
	router.HandleFunc("/transport/get_order", order_management.GetOrderHandler(orderManagementService)).Methods("GET")
	router.HandleFunc("/driver/set_location", driver.SetLocationHandler(driverService, &allocationService)).Methods("PUT")

	return router
}
