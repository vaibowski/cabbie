package driver

import (
	"cabbie/models"
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
)

type service interface {
	CreateNewDriver(driver models.Driver) (string, error)
}

func SignUpHandler(svc service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		reqBody, err := io.ReadAll(r.Body)
		if err != nil {
			log.Printf("error reading body: %s", err)
			handleError(w, errors.New("error reading request body"), http.StatusBadRequest)
			return
		}
		var signUpRequest SignUpRequest
		err = json.Unmarshal(reqBody, &signUpRequest)
		if err != nil {
			log.Printf("error unmarshalling body: %s", err)
			handleError(w, errors.New("error unmarshalling request body"), http.StatusBadRequest)
			return
		}
		driverID, err := svc.CreateNewDriver(models.Driver{
			Name:     signUpRequest.Name,
			Phone:    signUpRequest.Phone,
			Email:    signUpRequest.Email,
			Password: signUpRequest.Password,
		})
		if err != nil {
			log.Printf("error creating driver: %s", err)
			if err.Error() == "driver already exists" {
				handleError(w, err, http.StatusConflict)
				return

			} else {
				handleError(w, err, http.StatusInternalServerError)
				return
			}
		}
		json.NewEncoder(w).Encode(map[string]string{"driverID": driverID})
		return
	}
}

func handleError(w http.ResponseWriter, err error, code int) {
	http.Error(w, err.Error(), code)
}

type SignUpRequest struct {
	Name     string `json:"name"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
