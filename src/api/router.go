package api

import (
	"fmt"
	"time"
	"wedding/src/database"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/encryptcookie"
	"github.com/google/uuid"
)

func init() {
	App = fiber.New(fiber.Config{
		Immutable:       true,
		AppName:         "wedding",
		ReadTimeout:     10 * time.Millisecond,
		ReadBufferSize:  1024,
		RequestMethods:  []string{"GET", "POST", "HEAD", "DELETE"},
		ServerHeader:    "you are a curious dolphin",
		WriteTimeout:    10 * time.Millisecond,
		WriteBufferSize: 1024,
	})

	// Define CORS options
	corsConfig := cors.New(cors.Config{
		AllowOrigins: "http://localhost:3000", // Allow requests from localhost:3000
		AllowMethods: "GET,POST,PUT,DELETE",   // Allow specified HTTP methods
		AllowHeaders: "*",                     // Allow any headers
	})

	// Use CORS middleware with the specified options
	App.Use(corsConfig)

	App.Use("/guest/:uuid", func(c *fiber.Ctx) error {
		// Retrieve UUID from query parameter
		uuidString := c.Params("uuid")
		uuid, err := uuid.Parse(uuidString)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).
				JSON(fiber.Map{"errorMessage": err.Error()})
		}

		if !database.GuestExistsByUUID(uuid) {
			return c.Status(fiber.StatusNotFound).
				JSON(fiber.Map{"errorMessage": fmt.Errorf("user not found").Error()})
		}
		return c.Next()
	})
	COOKIEPASSWORD := readCookiePassword()
	App.Use(encryptcookie.New(encryptcookie.Config{
		Key:    COOKIEPASSWORD,
		Except: []string{"confirmed", "auth"},
	}))
	COOKIEPASSWORD = ""

	// login
	App.Post("/chisono", handleSecret)

	// landing page for QR code scan: LOGIN CSRF by default
	// there's no other way to make it easy for uncles ;)

	App.Get("/guest/:uuid", handleQRLoginGet)

	// authenticated routes
	api := App.Group("/api", authMiddleware)
	api.Get("/guest", handleFormGet)
	api.Post("/guest", handleFormPost)
	api.Delete("/guest", handleDelete)
	api.Get("/qr", handleQRGet)
	api.Get("/confirmed", handleConfirmedGet)
}
