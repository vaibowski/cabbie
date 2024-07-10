package customer

import (
	"cabbie/models"
	"encoding/json"
	"errors"
	"fmt"
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
			handleError(w, errors.New("error reading request body"), http.StatusBadRequest)
			return
		}

		err = json.Unmarshal(reqBody, &signUpRequest)
		if err != nil {
			log.Printf("error unmarshalling request body: %s", err)
			handleError(w, errors.New("error unmarshalling request body"), http.StatusBadRequest)
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
			if err.Error() == "customer already exists" {
				handleError(w, errors.New(fmt.Sprintf("error during signup: %s", err)), http.StatusConflict)
			} else {
				handleError(w, errors.New(fmt.Sprintf("error during signup: %s", err)), http.StatusInternalServerError)
			}
			return
		}
		json.NewEncoder(w).Encode(map[string]string{"customerID": customerID})
		return
	}
}

func handleError(w http.ResponseWriter, err error, code int) {
	http.Error(w, err.Error(), code)
}

type SignUpRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
}
