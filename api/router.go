package api

import (
	"net/http"

	"github.com/gorilla/mux"
)

var Router *mux.Router

func init() {
	// Create a new Gorilla Mux router
	Router = mux.NewRouter()

	// Define routes
	Router.HandleFunc("/form", handleForm).Methods("GET")
	Router.HandleFunc("/confirm", handleConfirmation).Methods("POST")

	// Serve static files (like CSS or JavaScript)
	Router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
}
