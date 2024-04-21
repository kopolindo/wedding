package api

import (
	"html/template"
	"net/http"
	"os"
	"wedding/database"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func Increment(x int) int {
	return x + 1
}

// HandleForm renders the form page
// method GET
// route /:uuid
func handleFormGet(c *fiber.Ctx) error {
	// Retrieve UUID from query parameter
	uuidString := c.Params("uuid")
	uuid, err := uuid.Parse(uuidString)
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
	// Execute the template with the data
	err = t.Execute(c.Response().BodyWriter(), guestMapSlice)
	if err != nil {
		return err
	}
	c.Type("html")
	return nil
}
