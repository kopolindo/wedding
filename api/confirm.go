package api

import (
	"log"
	"net/http"
	"strconv"
	"wedding/database"
	"wedding/models"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// handleFormPost handles form submission
func handleFormPost(c *fiber.Ctx) error {
	// Retrieve form values
	uuidString := c.Params("uuid")
	numberOfGuests := c.FormValue("guests")
	firstName := c.FormValue("firstname")
	lastName := c.FormValue("lastname")
	notes := []byte(c.FormValue("notes"))

	// Perform any necessary processing here
	// Cast number of guests
	numberOfGuestsParsed, err := strconv.Atoi(numberOfGuests)
	if err != nil {
		log.Printf("Error during conversion of number of guests. %s\n", err.Error())
	}

	// Parse UUID string to create uuid.UUID object
	uuidParsed, err := uuid.Parse(uuidString)
	if err != nil {
		log.Printf("Error during UUID parsing. %s\n", err.Error())
	}

	// Create guest struct from parsed data
	guest := models.Guest{
		FirstName:            firstName,
		LastName:             lastName,
		UUID:                 uuidParsed,
		NumberOfPartecipants: numberOfGuestsParsed,
		Confirmed:            true,
		Notes:                notes,
	}

	err = database.UpdateGuest(guest)
	if err != nil {
		log.Printf("error after updating guest: %s\n", err.Error())
		return c.JSON(fiber.Map{
			"errorMessage": err.Error(),
			"statusCode":   http.StatusInternalServerError,
		})
	}

	// Redirect to confirmation page
	return c.Redirect("/confirmation")
}
