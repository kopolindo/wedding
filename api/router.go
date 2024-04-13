package api

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

var Router http.Handler

func init() {
	// Create a new Gorilla Mux router
	router := mux.NewRouter()

	// Define routes
	router.HandleFunc("/form", handleForm).Methods("GET")
	router.HandleFunc("/confirm", handleConfirmation).Methods("POST")

	// Serve static files (like CSS or JavaScript)
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// Add timeout
	Router = http.TimeoutHandler(router, time.Second*3, "Timeout!")
}
