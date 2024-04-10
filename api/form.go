package api

import (
	"net/http"
	"text/template"
	"wedding/frontend"
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
	tmpl, err := template.New("form").Parse(frontend.FormTemplate)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
