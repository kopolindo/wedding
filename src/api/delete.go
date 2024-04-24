package api

import (
	"errors"
	"wedding/src/database"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type GuestToDelete struct {
	ID uint `json:"ID"`
}

// handleDelete removes guest from db
// method DELETE
// route /:uuid
func handleDelete(c *fiber.Ctx) error {
	// Retrieve UUID from query parameter
	uuidString := c.Params("uuid")
	uuid, err := uuid.Parse(uuidString)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).
			JSON(fiber.Map{"errorMessage": err.Error()})
	}
	guest := new(GuestToDelete)
	if err := c.BodyParser(guest); err != nil {
		return c.Status(fiber.StatusInternalServerError).
			JSON(fiber.Map{"errorMessage": err.Error()})
	}
	err = database.DeleteGuest(guest.ID, uuid)
	if err != nil {
		var e *fiber.Error
		if errors.As(err, &e) {
			return c.Status(e.Code).
				JSON(fiber.Map{"errorMessage": e.Message})
		}
	}
	return c.Status(fiber.StatusOK).
		JSON(fiber.Map{"message": "user successfully deleted"})
}
