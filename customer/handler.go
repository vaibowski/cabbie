package customer

import (
	"cabbie/models"
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
)

type service interface {
	CreateNewCustomer(customer models.Customer) (string, error)
}

func SignUpHandler(service service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("inside customer signup handler")
		var signUpRequest SignUpRequest
		reqBody, err := io.ReadAll(r.Body)
		if err != nil {
			log.Printf("error reading request body: %s", err)
			w.WriteHeader(http.StatusBadRequest)
			handleError(w, errors.New("error reading request body"))
			return
		}

		err = json.Unmarshal(reqBody, &signUpRequest)
		if err != nil {
			log.Printf("error unmarshalling request body: %s", err)
			w.WriteHeader(http.StatusBadRequest)
			handleError(w, errors.New("error unmarshalling request body"))
			return
		}

		// TODO: add request validation logic
		customer := models.Customer{
			Name:     signUpRequest.Name,
			Email:    signUpRequest.Email,
			Password: signUpRequest.Password,
			Phone:    signUpRequest.Phone,
		}
		customerID, err := service.CreateNewCustomer(customer)
		if err != nil {
			log.Printf("error during signup: %s", err)
			w.WriteHeader(http.StatusInternalServerError)
			handleError(w, errors.New("error during signup"))
			return
		}
		handleSuccess(w, customerID)
	}
}

type SignUpRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
}

func handleError(w http.ResponseWriter, err error) {
	http.Error(w, err.Error(), http.StatusInternalServerError)
}

func handleSuccess(w http.ResponseWriter, customerID string) {
	json.NewEncoder(w).Encode(map[string]string{"id": customerID})
	return
}
