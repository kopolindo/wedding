package main

import (
	"net/http"
	"wedding/api"
)

func main() {
	// Start the server
	http.ListenAndServe(":8080", api.Router)
}
