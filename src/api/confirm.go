package api

import (
	"log"
	"net/http"
	"wedding/src/database"
	"wedding/src/models"

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
	sessionID := c.Cookies("session")
	if sessionID == "" {
		return c.Status(fiber.StatusInternalServerError).
			JSON(fiber.Map{"errorMessage": "Hai trovato il modo di superare il cane a tre teste! Ma non supererai me..."})
	}
	log.Println(sessionID)
	uuid, err := uuid.Parse(sessionID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).
			JSON(fiber.Map{"errorMessage": err.Error()})
	}

	data := new(JSONForm)
	if err := c.BodyParser(data); err != nil {
		log.Printf("error during JSON parsing: %s\n", err.Error())
		return c.Status(fiber.StatusInternalServerError).
			JSON(fiber.Map{"errorMessage": err.Error()})
	}

	var guests []models.Guest
	// Create guest struct from parsed data
	for _, g := range data.People {
		guest := models.Guest{
			ID:        g.ID,
			FirstName: g.FirstName,
			LastName:  g.LastName,
			UUID:      uuid,
			Confirmed: true,
			Notes:     []byte(g.Notes),
		}
		guests = append(guests, guest)
	}

	for _, guest := range guests {
		if database.GuestExists(guest.ID, guest.UUID) {
			err = database.UpdateGuest(guest)
			if err != nil {
				log.Printf("error after updating guest: %s\n", err.Error())
				return c.JSON(fiber.Map{
					"errorMessage": err.Error(),
					"statusCode":   http.StatusInternalServerError,
				})
			}
		} else {
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
	return c.Status(fiber.StatusOK).
		JSON(fiber.Map{"status": "ok"})
}
