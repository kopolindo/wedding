package main

import (
	"net/http"
	"wedding/api"

	"github.com/gorilla/mux"
)

func main() {
	// Create a new Gorilla Mux router
	router := mux.NewRouter()

	// Define routes
	router.HandleFunc("/form", api.HandleForm).Methods("GET")
	router.HandleFunc("/confirm", api.HandleConfirmation).Methods("POST")

	// Serve static files (like CSS or JavaScript)
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// Start the server
	http.ListenAndServe(":8080", router)
}
