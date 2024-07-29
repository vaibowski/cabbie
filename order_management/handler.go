package order_management

import (
	"cabbie/models"
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
)

type service interface {
	CreateNewOrder(order models.Order) (models.Order, error)
}

func CreateOrderHandler(orderService service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()

		req, err := io.ReadAll(r.Body)
		if err != nil {
			log.Printf("error reading request body: %v", err)
			handleError(w, errors.New("error reading request body"), http.StatusBadRequest)
			return
		}

		var order models.Order
		err = json.Unmarshal(req, &order)
		if err != nil {
			log.Printf("error unmarshalling request body: %v", err)
			handleError(w, errors.New("error unmarshalling request body"), http.StatusBadRequest)
			return
		}
		order, err = orderService.CreateNewOrder(order)
		if err != nil {
			log.Printf("error creating new order: %v", err)
			handleError(w, errors.New("error creating new order"), http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(order)
	}
}

func handleError(w http.ResponseWriter, err error, code int) {
	http.Error(w, err.Error(), code)
}

type CreateOrderRequest struct {
	Origin      models.Location `json:"origin"`
	Destination models.Location `json:"destination"`
	ServiceType string          `json:"serviceType"`
	CustomerID  string          `json:"customerID"`
	TotalPrice  float64         `json:"totalPrice"`
}
