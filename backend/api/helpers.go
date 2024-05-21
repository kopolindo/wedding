package api

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
	"wedding/database"
	"wedding/log"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// readCookiePassword initializes COOKIEPASSWORD variable with content of COOKIEPASSWORDFILE
func readCookiePassword() string {
	if runningInDocker() {
		content, err := os.ReadFile(COOKIEPASSWORDFILEDOCKER)
		if err != nil {
			log.Errorf("Error reading password file. %s\n", err.Error())
			return ""
		}
		return string(content)
	} else {
		content, err := os.ReadFile(COOKIEPASSWORDFILE)
		if err != nil {
			log.Errorf("Error reading password file. %s\n", err.Error())
			return ""
		}
		return string(content)
	}
}

func authMiddleware(c *fiber.Ctx) error {
	c.Locals("isAuthenticated", false)
	// Check if the session cookie exists
	sessionID := c.Cookies("session")
	if sessionID != "" {
		u, err := uuid.Parse(sessionID)
		if err != nil {
			return c.Status(fiber.StatusForbidden).
				JSON(fiber.Map{"errorMessage": fmt.Sprintf("RON SI E' SPACCATO! %s", err.Error())})
		}
		if database.GuestExistsByUUID(u) {
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

// confirmedCookie function check if guest confirmated and set confirmed cookie
func confirmedCookie(uuid uuid.UUID) (fiber.Cookie, error) {
	// this cookie is not really important for authentication purpose
	// it's only required for client-side rendering of "qr" page
	// HTTPOnly is false because client-side code needs to check the presence
	// of confirmed cookie to render the page.
	guests, err := database.GetUsersByUUID(uuid)
	if err != nil {
		return fiber.Cookie{}, err
	}
	confirmed := "false"
	if guests[0].Confirmed {
		confirmed = "true"
	}
	return fiber.Cookie{
		Name:     "confirmed",
		Value:    confirmed,
		Expires:  time.Now().Add(24 * 7 * time.Hour),
		Secure:   true,
		HTTPOnly: false,
		SameSite: "strict",
	}, nil
}

// runningInDocker checks if application is running in docker container
// returns true if running in docker, false otherwise
func runningInDocker() bool {
	file, err := os.Open("/proc/1/cgroup")
	if err != nil {
		log.Errorf("Error: %s", err.Error())
		return false
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "docker") || strings.Contains(line, "docker-") {
			return true
		}
	}

	return false
}
