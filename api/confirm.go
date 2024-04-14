package api

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"wedding/database"
	"wedding/models"

	"github.com/google/uuid"
)

// HandleConfirmation handles form submission
func handleConfirmation(w http.ResponseWriter, r *http.Request) {
	// Retrieve form values
	err := r.ParseForm()
	if err != nil {
		log.Printf("error during request parsing: %s\n", err.Error())
	}
	numberOfGuests := r.Form.Get("guests")
	firstName := r.Form.Get("firstname")
	lastName := r.Form.Get("lastname")
	uuidValue := r.Form.Get("uuid")

	numberOfGuestsParsed, err := strconv.Atoi(numberOfGuests)
	if err != nil {
		log.Printf("Error during conversion of number of guests. %s\n", err.Error())
	}

	// Perform any necessary processing here
	// For demonstration, we'll just print the values
	println("UUID:", uuidValue)
	println("Guests:", numberOfGuests)
	println("First name:", firstName)
	println("Last name:", lastName)

	uuidParsed, err := uuid.Parse(uuidValue)
	if err != nil {
		log.Printf("Error during UUID parsing. %s\n", err.Error())
	}

	guest := models.Guest{
		FirstName:            firstName,
		LastName:             lastName,
		UUID:                 uuidParsed,
		NumberOfPartecipants: numberOfGuestsParsed,
		Confirmed:            true,
		Notes:                []byte{},
	}

	index, err := database.InsertGuestData(guest)
	if err != nil {
		log.Printf("Error during database insert. %s\n", err.Error())
	} else {
		fmt.Println(index)
	}

	// Return a success message
	_, err = w.Write([]byte("Confirmation received!"))
	if err != nil {
		log.Printf("error during reposonse writing: %s\n", err.Error())
	}
}
