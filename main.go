package main

import (
	"net/http"
	"wedding/api"

	"github.com/gorilla/mux"
)

// HandleConfirmation handles form submission
func HandleConfirmation(w http.ResponseWriter, r *http.Request) {
	// Retrieve form values
	r.ParseForm()
	guests := r.Form.Get("guests")
	uuid := r.Form.Get("uuid")

	// Perform any necessary processing here
	// For demonstration, we'll just print the values
	println("UUID:", uuid)
	println("Guests:", guests)

	// You can add further logic here like saving to a database, etc.

	// Return a success message
	w.Write([]byte("Confirmation received!"))
}

func main() {
	// Create a new Gorilla Mux router
	router := mux.NewRouter()

	// Define routes
	router.HandleFunc("/form", api.HandleForm).Methods("GET")
	router.HandleFunc("/confirm", HandleConfirmation).Methods("POST")

	// Serve static files (like CSS or JavaScript)
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// Start the server
	http.ListenAndServe(":8080", router)
}
