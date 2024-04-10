package main

import (
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
)

// PageData represents data to be rendered in HTML template
type PageData struct {
	UUID string
}

// HandleForm renders the form page
func HandleForm(w http.ResponseWriter, r *http.Request) {
	// Retrieve UUID from query parameter
	uuid := r.URL.Query().Get("uuid")

	// Pass UUID to the template
	data := PageData{UUID: uuid}

	// Render the HTML template
	tmpl, err := template.New("form").Parse(formTemplate)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

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
	router.HandleFunc("/form", HandleForm).Methods("GET")
	router.HandleFunc("/confirm", HandleConfirmation).Methods("POST")

	// Serve static files (like CSS or JavaScript)
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// Start the server
	http.ListenAndServe(":8080", router)
}

// HTML template for the form
var formTemplate = `
<!DOCTYPE html>
<html>
<head>
    <title>Confirmation Form</title>
    <link rel="stylesheet" type="text/css" href="/static/style.css">
</head>
<body>
    <h1>Confirmation Form</h1>
    <form action="/confirm" method="post">
        <input type="hidden" id="uuid" name="uuid">
        <label for="guests">Number of guests:</label>
        <input type="number" id="guests" name="guests" required>
        <br>
        <input type="submit" value="Confirm">
    </form>	
	<script>
	// Function to extract query parameter from URL
	function getQueryParam(name) {
		const urlParams = new URLSearchParams(window.location.search);
		return urlParams.get(name);
	}

	// Set the UUID value from query string to the hidden input field
	const uuidInput = document.getElementById('uuid');
	const uuidValue = getQueryParam('uuid');
	uuidInput.value = uuidValue;

	// Submit the form
	document.getElementById('confirmationForm').onsubmit = function() {
		if (!uuidValue) {
			alert('UUID not found in query string');
			return false; // Prevent form submission if UUID is missing
		}
	};
</script>
</body>
</html>
`
