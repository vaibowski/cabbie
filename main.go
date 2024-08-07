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
	"time"
)

func main() {
	customerDB := make(map[string]models.Customer)
	initializeCustomers(customerDB)
	driverDB := make(map[string]models.Driver)
	orderDB := make(map[string]models.Order)
	var activeDriverPool []*treemap.Map[float64, []string]
	for i := 0; i <= 4; i++ {
		m := treemap.New[float64, []string]()
		activeDriverPool = append(activeDriverPool, m)
	}
	initializeDrivers(driverDB, activeDriverPool)

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

func initializeCustomers(customerDB map[string]models.Customer) {
	for i := 0; i < 4; i++ {
		customerID := fmt.Sprintf("customer_%d", i)
		customerName := fmt.Sprintf("vaibhav_%d", i)
		email := "email@email.com"
		password := "password"
		phone := fmt.Sprintf("phone_%d", i)
		customer := models.Customer{
			CustomerID: customerID,
			Name:       customerName,
			Email:      email,
			Password:   password,
			Phone:      phone,
			CreatedAt:  time.Time{},
		}
		customerDB[customerID] = customer
		customerDB[phone] = customer
	}
}

func initializeDrivers(driverDB map[string]models.Driver, activeDriverPool []*treemap.Map[float64, []string]) {
	for i := 1; i <= 4; i++ {
		driverID := fmt.Sprintf("driver_%d", i)
		driverName := fmt.Sprintf("driver_%d", i)
		email := "driveremail@email.com"
		password := "password"
		phone := fmt.Sprintf("phone_%d", i)
		serviceType := models.ServiceTypeEnum(i)
		driver := models.Driver{
			DriverID:     driverID,
			Name:         driverName,
			Email:        email,
			Password:     password,
			Phone:        phone,
			ServiceType:  serviceType,
			LastLocation: models.Location{XCoordinate: float64(i * 10)},
			CreatedAt:    time.Time{},
		}
		driverDB[driverID] = driver
		driverDB[phone] = driver

		activeDriverPool[serviceType].Put(float64(i*10), []string{driverID})
	}
}
