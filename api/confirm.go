package api

import "net/http"

// HandleConfirmation handles form submission
func handleConfirmation(w http.ResponseWriter, r *http.Request) {
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
