package main

import (
	"errors"
	"fmt"
	"io"
	logger "log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/ping", pingHandler).Methods("GET")
	logger.Println("Listening on port 8080")
	err := http.ListenAndServe(":8080", router)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}

func pingHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "pong")
	return
}
