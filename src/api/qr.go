package api

import (
	"encoding/json"
	"fmt"
	"wedding/src/database"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// handleQrGet renders the qr generation
// method POST
// route /api/qr
func handleQRPost(c *fiber.Ctx) error {
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

	if !database.GuestExistsByUUID(uuid) {
		return c.Status(fiber.StatusNotFound).
			JSON(fiber.Map{"errorMessage": "user not found...who you really are??"})
	}

	payload := &qrRequestPayload{}
	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusNotFound).
			JSON(fiber.Map{"errorMessage": err.Error()})
	}
	id := payload.ID

	guest, err := database.GetUserByID(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).
			JSON(fiber.Map{"errorMessage": err.Error()})
	}

	if guest.UUID != uuid {
		return c.Status(fiber.StatusNotFound).
			JSON(fiber.Map{"errorMessage": "non penso proprio ragazzino"})
	}

	uuidString := guest.UUID.String()

	response := &qrResponseBody{
		Success: true,
		Message: fmt.Sprintf("qr correctly generated [%s]", uuidString),
	}

	responseJSON, err := json.Marshal(response)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).
			JSON(fiber.Map{"errorMessage": err.Error()})
	}
	c.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)
	return c.Send(responseJSON)
}
