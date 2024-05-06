package api

import (
	"errors"
	"time"
	"wedding/src/database"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// handleDelete removes guest from db
// method GET
// route /api/confirmed
func handleConfirmedGet(c *fiber.Ctx) error {
	// Retrieve UUID from query parameter
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
		var e *fiber.Error
		if errors.As(err, &e) {
			return c.Status(e.Code).
				JSON(fiber.Map{"errorMessage": e.Message})
		}
	}
	if guests[0].Confirmed {
		// this cookie is not really important for authentication purpose
		// it's only required for client-side rendering of "qr" page
		// HTTPOnly is false because client-side code needs to check the presence
		// of confirmed cookie to render the page.
		c.Cookie(&fiber.Cookie{
			Name:     "confirmed",
			Value:    "true",
			Expires:  time.Now().Add(24 * 7 * time.Hour),
			Secure:   true,
			HTTPOnly: false,
			SameSite: "strict",
		})
	}
	return c.SendStatus(fiber.StatusOK)
}
