package api

import (
	"fmt"
	"log"
	"net/http"
	"time"
	"wedding/src/database"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/encryptcookie"
	"github.com/gofiber/template/html/v2"
	"github.com/google/uuid"
)

var (
	Router http.Handler
	App    *fiber.App
)

func init() {
	engine := html.New("./views", ".html")
	App = fiber.New(fiber.Config{
		Immutable:       true,
		AppName:         "wedding",
		ReadTimeout:     10 * time.Millisecond,
		ReadBufferSize:  1024,
		RequestMethods:  []string{"GET", "POST", "HEAD", "DELETE"},
		ServerHeader:    "you are a curious dolphin",
		Views:           engine,
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
		Key: COOKIEPASSWORD,
	}))
	COOKIEPASSWORD = ""

	App.Use(func(c *fiber.Ctx) error {
		c.Locals("isAuthenticated", false)
		// Check if the session cookie exists
		sessionID := c.Cookies("session")
		if sessionID != "" {
			log.Printf("found session cookie: %s\n", sessionID)
			u, err := uuid.Parse(sessionID)
			if err != nil {
				return c.Status(fiber.StatusForbidden).
					JSON(fiber.Map{"errorMessage": fmt.Sprintf("you are not authenticated: %s", err.Error())})
			}
			if database.GuestExistsByUUID(u) {
				log.Println("UUID is correctly formatted and a user with such UUID exists")
				c.Locals("isAuthenticated", true)
			} else {
				return c.Status(fiber.StatusForbidden).
					JSON(fiber.Map{"errorMessage": fmt.Sprintln("you are not authenticated")})
			}
		}
		return c.Next()
	})
	App.Get("/guest/:uuid", handleFormGet)
	App.Post("/guest/:uuid", handleFormPost)
	App.Delete("/guest/:uuid", handleDelete)
	App.Get("/confirmation", func(c *fiber.Ctx) error {
		return c.SendFile("./views/confirmation.html")
	})
	App.Get("/chisono", func(c *fiber.Ctx) error {
		return c.SendFile("./views/secret.html")
	})
	App.Post("/chisono", handleSecret)
}
