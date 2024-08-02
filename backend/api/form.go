package api

import (
	"encoding/json"
	"wedding/database"
	"wedding/log"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// HandleForm renders the form page
// method GET
// route /api/guest
func handleFormGet(c *fiber.Ctx) error {
	sessionID := c.Cookies("session")
	if sessionID == "" {
		return c.Status(fiber.StatusInternalServerError).
			JSON(fiber.Map{"errorMessage": "Hai trovato il modo di superare il cane a tre teste! Ma non supererai me..."})
	}
	uuid, err := uuid.Parse(sessionID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).
			JSON(fiber.Map{"errorMessage": err.Error()})
	}

	guests, err := database.GetUsersByUUID(uuid)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).
			JSON(fiber.Map{"errorMessage": err.Error()})
	}

	firstName, lastName, err := database.GetMainUserByUUID(uuid)
	if err != nil {
		log.Errorf("error while getting main user by uuid (%s): %s", uuid, err.Error())
	}
	log.Debugf("hi user %s %s (%s)\n", firstName, lastName, uuid)

	//var guestMapSlice []map[string]interface{}
	var respGuests []responseGuest
	for _, guest := range guests {
		respGuest := responseGuest{
			ID:        uint(guest.ID),
			FirstName: guest.FirstName,
			LastName:  guest.LastName,
			Confirmed: guest.Confirmed,
			Notes:     guest.Notes,
			Type:      guest.Type,
		}
		respGuests = append(respGuests, respGuest)
	}
	responseGuests := responseGuests{
		Guests: respGuests,
	}
	guestsJSON, err := json.Marshal(responseGuests)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).
			JSON(fiber.Map{"errorMessage": err.Error()})
	}
	cookie, err := confirmedCookie(uuid)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).
			JSON(fiber.Map{"errorMessage": err.Error()})
	}
	c.Cookie(&cookie)
	c.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)
	return c.Send(guestsJSON)
}
