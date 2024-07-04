package main

import (
	"cabbie/customer"
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
	customerRepo := repository.CustomerRepository{MapDatastore: customerDB}
	customerService := customer.NewService(&customerRepo)
	router := NewRouter(customerService)

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
