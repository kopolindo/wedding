package api

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"wedding/database"

	"github.com/google/uuid"
)

// HandleConfirmation handles form submission
func handleConfirmation(w http.ResponseWriter, r *http.Request) {
	// Retrieve form values
	r.ParseForm()
	numberOfGuests := r.Form.Get("guests")
	uuidValue := r.Form.Get("uuid")

	numberOfGuestsParsed, err := strconv.Atoi(numberOfGuests)
	if err != nil {
		log.Printf("Error during conversion of number of guests. %s\n", err.Error())
	}

	// Perform any necessary processing here
	// For demonstration, we'll just print the values
	println("UUID:", uuidValue)
	println("Guests:", numberOfGuests)

	uuidParsed, err := uuid.Parse(uuidValue)
	if err != nil {
		log.Printf("Error during UUID parsing. %s\n", err.Error())
	}

	guest := database.Guest{
		UUID:                 uuidParsed,
		NumberOfPartecipants: numberOfGuestsParsed,
	}

	index, err := database.InsertGuestData(guest)
	if err != nil {
		log.Printf("Error during database insert. %s\n", err.Error())
	} else {
		fmt.Println(index)
	}

	// Return a success message
	w.Write([]byte("Confirmation received!"))
}
