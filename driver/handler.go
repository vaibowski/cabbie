package driver

import (
	"cabbie/models"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
)

type driverService interface {
	CreateNewDriver(driver models.Driver) (string, error)
	FetchDriver(driverID string) (models.Driver, error)
	UpdateDriver(driver models.Driver) error
}

type allocationService interface {
	SetLocation(driverID string, serviceType models.ServiceTypeEnum, location models.Location)
	UnsetLocation(driverID string, serviceType models.ServiceTypeEnum, lastLocation models.Location) error
}

func SignUpHandler(svc driverService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
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

func SetLocationHandler(driverService driverService, allocationService allocationService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()

		reqBody, err := io.ReadAll(r.Body)
		if err != nil {
			log.Printf("error reading request body: %s", err.Error())
			handleError(w, errors.New("error reading request body"), http.StatusBadRequest)
			return
		}

		var request setLocationRequest
		err = json.Unmarshal(reqBody, &request)
		if err != nil {
			log.Printf("error unmarshalling request body: %s", err.Error())
			return
		}

		// fetch driver
		driver, err := driverService.FetchDriver(request.DriverID)
		if err != nil {
			log.Printf("error fetching driver: %s", err.Error())
			handleError(w, err, http.StatusNotFound)
			return
		}

		// unset last location
		err = allocationService.UnsetLocation(driver.DriverID, driver.ServiceType, request.Location)
		if err != nil {
			log.Printf("error unsetting last location: %s", err.Error())
			handleError(w, err, http.StatusInternalServerError)
		}

		// update/set location as per request
		allocationService.SetLocation(driver.DriverID, driver.ServiceType, request.Location)

		// update driver's last location and update in driver repository
		driver.LastLocation = request.Location
		err = driverService.UpdateDriver(driver)
		if err != nil {
			// TODO: need to rollback driver set location if update driver fails
			log.Printf("error updating driver: %s", err.Error())
			handleError(w, err, http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(map[string]string{"driverID": driver.DriverID, "location": fmt.Sprintf("%f", driver.LastLocation)})
		return
	}
}

func handleError(w http.ResponseWriter, err error, code int) {
	http.Error(w, err.Error(), code)
}

type SignUpRequest struct {
	Name        string `json:"name"`
	Phone       string `json:"phone"`
	Email       string `json:"email"`
	ServiceType string `json:"service_type"`
	Password    string `json:"password"`
}

type setLocationRequest struct {
	DriverID string          `json:"driverID"`
	Location models.Location `json:"location"`
}
