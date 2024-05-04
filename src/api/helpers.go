package api

import (
	"fmt"
	"log"
	"os"
	"wedding/src/database"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// readCookiePassword initializes COOKIEPASSWORD variable with content of COOKIEPASSWORDFILE
func readCookiePassword() string {
	content, err := os.ReadFile(COOKIEPASSWORDFILE)
	if err != nil {
		log.Fatalf("Error reading password file. %s\n", err.Error())
		return ""
	}
	return string(content)
}

func authMiddleware(c *fiber.Ctx) error {
	c.Locals("isAuthenticated", false)
	// Check if the session cookie exists
	sessionID := c.Cookies("session")
	if sessionID != "" {
		log.Printf("found session cookie: %s\n", sessionID)
		u, err := uuid.Parse(sessionID)
		if err != nil {
			return c.Status(fiber.StatusForbidden).
				JSON(fiber.Map{"errorMessage": fmt.Sprintf("RON SI E' SPACCATO! %s", err.Error())})
		}
		if database.GuestExistsByUUID(u) {
			log.Println("UUID is correctly formatted and a user with such UUID exists")
			c.Locals("isAuthenticated", true)
		} else {
			return c.Status(fiber.StatusForbidden).
				JSON(fiber.Map{"errorMessage": "Tu osi usare i miei incantesimi contro di me???"})
		}
	} else {
		return c.Status(fiber.StatusForbidden).
			JSON(fiber.Map{"errorMessage": "Niente permesso, niente gita al villaggio Potter!"})
	}
	return c.Next()
}
