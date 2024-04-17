package api

import (
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
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
		RequestMethods:  []string{"GET", "POST", "HEAD"},
		ServerHeader:    "you are a curious dolphin",
		Views:           engine,
		WriteTimeout:    10 * time.Millisecond,
		WriteBufferSize: 1024,
	})

	//App.Static("/", "../static/")
	App.Get("/guest/:uuid", handleFormGet)
	App.Post("/guest/:uuid", handleFormPost)
	App.Get("/confirmation", func(c *fiber.Ctx) error {
		return c.SendFile("./views/confirmation.html")
	})

}
