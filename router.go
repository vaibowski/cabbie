package main

import (
	"cabbie/customer"
	"github.com/gorilla/mux"
)

func NewRouter(customerService customer.Service) *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/ping", pingHandler).Methods("GET")
	router.HandleFunc("/customer/signup", customer.SignUpHandler(customerService)).Methods("POST")
	return router
}
