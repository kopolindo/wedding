package api

import (
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
	cookie, err := confirmedCookie(uuid)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).
			JSON(fiber.Map{"errorMessage": err.Error()})
	}
	c.Cookie(&cookie)
	return c.SendStatus(fiber.StatusOK)
}
