package api

import (
	"encoding/base64"
	"wedding/backend/database"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	qrcode "github.com/skip2/go-qrcode"
)

// handleQrGet renders the qr generation
// method GET
// route /api/qr
func handleQRGet(c *fiber.Ctx) error {
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
	guests, err := database.GetUsersByUUID(uuid)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).
			JSON(fiber.Map{"errorMessage": err.Error()})
	}
	if guests[0].Confirmed {
		var png []byte
		png, err = qrcode.Encode(uuid.String(), qrcode.Medium, 256)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).
				JSON(fiber.Map{"errorMessage": err.Error()})
		}
		sEnc := base64.StdEncoding.EncodeToString(png)
		c.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)
		return c.Status(fiber.StatusOK).
			JSON(fiber.Map{"qrcode": sEnc})
	} else {
		return c.Status(fiber.StatusForbidden).
			JSON(fiber.Map{"errorMessage": "Prima confermi, poi ti do un QR code ;)"})
	}
}
