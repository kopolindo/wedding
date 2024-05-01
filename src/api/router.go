package api

import (
	"fmt"
	"net/http"
	"time"
	"wedding/src/database"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/template/html/v2"
	"github.com/google/uuid"
)

var Router http.Handler
var App *fiber.App

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
