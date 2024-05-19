package api

import (
	"time"
	"wedding/src/database"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// handleQrGet renders the qr generation
// method GET
// route /:uuid
func handleQRLoginGet(c *fiber.Ctx) error {
	uuidString := c.Params("uuid")
	u, err := uuid.Parse(uuidString)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(fiber.Map{"errorMessage": err.Error()})
	}
	if !database.GuestExistsByUUID(u) {
		return c.Status(fiber.StatusNotFound).
			JSON(fiber.Map{"errorMessage": "user not found...who you really are??"})
	}

	c.Cookie(&fiber.Cookie{
		Name:     "session",
		Value:    uuidString,
		Expires:  time.Now().Add(24 * 7 * time.Hour),
		Secure:   true,
		HTTPOnly: true,
		SameSite: "strict",
	})
	// this second cookie is not really important for authentication purpose
	// it's only required for client-side rendering of "private" pages
	// in React nothing is "really" private.
	// HTTPOnly is false because client-side code needs to check the presence
	// of auth cookie to render pages. Even if users forge such cookie, it is
	// not used server-side to authenticate requests.
	c.Cookie(&fiber.Cookie{
		Name:     "auth",
		Value:    "true",
		Expires:  time.Now().Add(24 * 7 * time.Hour),
		Secure:   true,
		HTTPOnly: false,
		SameSite: "strict",
	})
	cookie, err := confirmedCookie(u)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).
			JSON(fiber.Map{"errorMessage": err.Error()})
	}
	c.Cookie(&cookie)
	return c.SendStatus(fiber.StatusOK)
}
