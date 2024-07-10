package estimate

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
)

type service interface {
	ServeEstimate(req Request) (Response, error)
}

func Handler(s service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		reqBody, err := io.ReadAll(r.Body)
		if err != nil {
			handleError(w, errors.New("failed to read request body"), http.StatusBadRequest)
			return
		}
		var estimateReq Request
		err = json.Unmarshal(reqBody, &estimateReq)
		if err != nil {
			log.Printf("failed to unmarshal request body: %s", err)
			handleError(w, errors.New("failed to unmarshal request body"), http.StatusBadRequest)
			return
		}
		response, err := s.ServeEstimate(estimateReq)
		if err != nil {
			log.Printf("failed to serve estimate: %s", err)
			handleError(w, errors.New("failed to fetch estimate"), http.StatusInternalServerError)
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}
}

func handleError(w http.ResponseWriter, err error, code int) {
	http.Error(w, err.Error(), code)
}
