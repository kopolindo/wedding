package api

import (
	"github.com/gofiber/fiber/v2"
)

// handleSecret renders the login page
// method POST
// route /secret
func handleSecret(c *fiber.Ctx) error {
	// Parse form data
	c.FormValue("secret")
	// Redirect to a success page or return a response
	c.Type("html")
	return nil
}
