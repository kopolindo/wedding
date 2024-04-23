package api

import (
	"net/http"
	"net/url"
	"wedding/backend"
	"wedding/database"

	"github.com/gofiber/fiber/v2"
)

// handleSecret renders the login page
// method POST
// route /secret
func handleSecret(c *fiber.Ctx) error {
	c.Type("html")
	// Parse form data
	secret := c.FormValue("secret")
	guests := database.GetAllUsers()
	for _, guest := range guests {
		ok, err := backend.ComparePasswordAndHash(secret, guest.Secret)
		if err != nil {
			return c.JSON(fiber.Map{
				"errorMessage": err.Error(),
				"statusCode":   http.StatusInternalServerError,
			})
		}
		if ok {
			location, err := url.JoinPath("/guest", guest.UUID.String())
			if err != nil {
				return c.JSON(fiber.Map{
					"errorMessage": err.Error(),
					"statusCode":   http.StatusInternalServerError,
				})
			}
			return c.Redirect(location)
		}
	}
	return c.Redirect("/chisono")
}
