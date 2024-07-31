package main

import (
	"cabbie/customer"
	"cabbie/driver"
	"cabbie/estimate"
	"cabbie/models"
	"cabbie/order_management"
	"cabbie/repository"
	"errors"
	"fmt"
	"github.com/emirpasic/gods/v2/maps/treemap"
	"io"
	logger "log"
	"net/http"
	"os"
)

func main() {
	customerDB := make(map[string]models.Customer)
	driverDB := make(map[string]models.Driver)
	orderDB := make(map[string]models.Order)
	var activeDriverPool []*treemap.Map[float64, []string]
	for i := 0; i <= 4; i++ {
		m := treemap.New[float64, []string]()
		activeDriverPool = append(activeDriverPool, m)
	}

	customerRepo := repository.CustomerRepository{MapDatastore: customerDB}
	driverRepo := repository.DriverRepository{MapDatastore: driverDB}
	orderRepo := repository.OrderRepository{MapDatastore: orderDB}

	customerService := customer.NewService(&customerRepo)
	driverService := driver.NewService(&driverRepo)
	estimateService := estimate.NewService()
	allocationService := driver.NewAllocationService(activeDriverPool)
	orderService := order_management.NewService(&orderRepo, &allocationService)
	router := NewRouter(customerService, driverService, estimateService, allocationService, orderService)

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
