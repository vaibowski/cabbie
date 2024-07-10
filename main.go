package main

import (
	"cabbie/customer"
	"cabbie/driver"
	"cabbie/estimate"
	"cabbie/models"
	"cabbie/repository"
	"errors"
	"fmt"
	"io"
	logger "log"
	"net/http"
	"os"
)

func main() {
	customerDB := make(map[string]models.Customer)
	driverDB := make(map[string]models.Driver)

	customerRepo := repository.CustomerRepository{MapDatastore: customerDB}
	driverRepo := repository.DriverRepository{MapDatastore: driverDB}

	customerService := customer.NewService(&customerRepo)
	driverService := driver.NewService(&driverRepo)
	estimateService := estimate.NewService()
	router := NewRouter(customerService, driverService, estimateService)

	logger.Println("Listening on port 8080")
	err := http.ListenAndServe(":8080", router)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}

func pingHandler(w http.ResponseWriter, _ *http.Request) {
	io.WriteString(w, "pong")
	return
}
