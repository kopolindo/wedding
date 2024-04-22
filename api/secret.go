package api

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

// handleSecret renders the login page
// method POST
// route /secret
func handleSecret(c *fiber.Ctx) error {
	// Parse form data
	secret := c.FormValue("secret")
	fmt.Println(secret)
	// Redirect to a success page or return a response
	c.Type("html")
	return nil
}
