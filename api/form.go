package api

import (
	"net/http"
	"wedding/database"
	"wedding/models"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// HandleForm renders the form page
// method GET
// route /:uuid
func handleFormGet(c *fiber.Ctx) error {
	// Retrieve UUID from query parameter
	uuidString := c.Params("uuid")
	uuid, err := uuid.Parse(uuidString)
	if err != nil {
		return c.JSON(fiber.Map{
			"errorMessage": err.Error(),
			"statusCode":   http.StatusInternalServerError,
		})
	}
	guest, err := database.GetUserByUUID(uuid)
	if err != nil {
		return c.JSON(fiber.Map{
			"errorMessage": err.Error(),
			"statusCode":   http.StatusInternalServerError,
		})
	}
	guestMap := models.StructToMap(guest)
	return c.Render("form", guestMap)
}
