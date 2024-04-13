package main

import (
	"log"
	"net/http"
	"wedding/api"
)

func main() {
	// Start the server
	err := http.ListenAndServe(":8080", api.Router)
	if err != nil {
		log.Fatalf("error starting the server: %s\n", err.Error())
	}
}
