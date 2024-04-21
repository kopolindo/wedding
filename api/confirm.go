package api

import (
	"fmt"
	"log"
	"net/http"
	"wedding/database"
	"wedding/models"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type JSONForm struct {
	Guests int `json:"guests"`
	People []struct {
		ID        uint   `json:"ID"`
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Notes     string `json:"notes"`
	} `json:"people"`
}

// handleFormPost handles form submission
func handleFormPost(c *fiber.Ctx) error {
	// Retrieve form values
	uuidString := c.Params("uuid")
	// Parse UUID string to create uuid.UUID object
	uuidParsed, err := uuid.Parse(uuidString)
	if err != nil {
		log.Printf("Error during UUID parsing. %s\n", err.Error())
	}

	data := new(JSONForm)
	if err := c.BodyParser(data); err != nil {
		log.Printf("error during JSON parsing: %s\n", err.Error())
		return c.JSON(fiber.Map{
			"errorMessage": err.Error(),
			"statusCode":   http.StatusInternalServerError,
		})
	}

	var guests []models.Guest
	// Create guest struct from parsed data
	for _, g := range data.People {
		guest := models.Guest{
			ID:        g.ID,
			FirstName: g.FirstName,
			LastName:  g.LastName,
			UUID:      uuidParsed,
			Confirmed: true,
			Notes:     []byte(g.Notes),
		}
		guests = append(guests, guest)
	}

	fmt.Println(guests)
	for _, guest := range guests {
		log.Printf("Check existence of guest %v\n", guest)
		if database.GuestExists(guest) {
			log.Printf("Updating guest: %v\n", guest)
			err = database.UpdateGuest(guest)
			if err != nil {
				log.Printf("error after updating guest: %s\n", err.Error())
				return c.JSON(fiber.Map{
					"errorMessage": err.Error(),
					"statusCode":   http.StatusInternalServerError,
				})
			}
		} else {
			log.Printf("Creating guest: %v\n", guest)
			_, err := database.CreateGuest(guest)
			if err != nil {
				log.Printf("error after updating guest: %s\n", err.Error())
				return c.JSON(fiber.Map{
					"errorMessage": err.Error(),
					"statusCode":   http.StatusInternalServerError,
				})
			}
		}
	}

	// Redirect to confirmation page
	return c.Redirect("/confirmation")
}
