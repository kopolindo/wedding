package api

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"os"
	"wedding/src/database"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func Increment(x int) int {
	return x + 1
}

// HandleForm renders the form page
// method GET
// route /guest
func handleFormGet(c *fiber.Ctx) error {
	sessionID := c.Cookies("session")
	log.Println(sessionID)
	uuid, err := uuid.Parse(sessionID)
	if err != nil {
		return c.JSON(fiber.Map{
			"errorMessage": err.Error(),
			"statusCode":   http.StatusInternalServerError,
		})
	}

	tmplFile := "./views/form.html"
	tmplContent, err := os.ReadFile(tmplFile)
	if err != nil {
		panic(err)
	}
	t := template.New("form").Funcs(template.FuncMap{"Increment": Increment})
	// Parse the template content
	_, err = t.Parse(string(tmplContent))
	if err != nil {
		panic(err)
	}

	guests, err := database.GetUsersByUUID(uuid)
	if err != nil {
		return c.JSON(fiber.Map{
			"errorMessage": err.Error(),
			"statusCode":   http.StatusInternalServerError,
		})
	}
	var guestMapSlice []map[string]interface{}
	for _, guest := range guests {
		guestMap := map[string]interface{}{
			"ID":        guest.ID,
			"FirstName": guest.FirstName,
			"LastName":  guest.LastName,
			"UUID":      guest.UUID,
			"Confirmed": guest.Confirmed,
			"Notes":     string(guest.Notes), // Convert []byte to string
		}
		guestMapSlice = append(guestMapSlice, guestMap)
	}
	// Marshal guestMapSlice into JSON
	guestsJSON, err := json.Marshal(guestMapSlice)
	if err != nil {
		return err
	}

	// Write JSON response
	c.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)
	return c.Send(guestsJSON)
}
