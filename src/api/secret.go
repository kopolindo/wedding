package api

import (
	"encoding/json"
	"log"
	"sync"
	"time"
	"wedding/src/backend"
	"wedding/src/database"
	"wedding/src/models"

	"github.com/gofiber/fiber/v2"
)

type Payload struct {
	Secret string `json:"secret"`
}

// handleSecret renders the login page
// method POST
// route /secret
func handleSecret(c *fiber.Ctx) error {
	start := time.Now()
	// Parse form data
	payload := &Payload{}
	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusNotFound).
			JSON(fiber.Map{"errorMessage": err.Error()})
	}
	secret := payload.Secret
	guests := database.GetAllUsers()
	type result struct {
		guest models.Guest
		ok    bool
		err   error
	}

	results := make(chan result, len(guests))
	var wg sync.WaitGroup
	for _, guest := range guests {
		wg.Add(1)
		go func(guest models.Guest) {
			defer wg.Done()
			ok, err := backend.ComparePasswordAndHash(secret, guest.Secret)
			results <- result{guest, ok, err}
		}(guest)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	for range guests {
		res := <-results
		if res.err != nil {
			return c.Status(fiber.StatusNotFound).
				JSON(fiber.Map{"errorMessage": res.err.Error()})
		}
		if res.ok {
			uuid := res.guest.UUID
			guests, err := database.GetUsersByUUID(uuid)
			if err != nil {
				return c.Status(fiber.StatusInternalServerError).
					JSON(fiber.Map{"errorMessage": err.Error()})
			}
			var guestMapSlice []map[string]interface{}
			for _, guest := range guests {
				guestMap := map[string]interface{}{
					"ID":        guest.ID,
					"FirstName": guest.FirstName,
					"LastName":  guest.LastName,
					"Confirmed": guest.Confirmed,
					"Notes":     string(guest.Notes), // Convert []byte to string
				}
				guestMapSlice = append(guestMapSlice, guestMap)
			}
			// Marshal guestMapSlice into JSON
			guestsJSON, err := json.Marshal(guestMapSlice)
			if err != nil {
				return c.Status(fiber.StatusInternalServerError).
					JSON(fiber.Map{"errorMessage": err.Error()})
			}
			// Write JSON response
			c.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)
			c.Cookie(&fiber.Cookie{
				Name:     "UUID",
				Value:    uuid.String(),
				Expires:  time.Now().Add(24 * 7 * time.Hour),
				Secure:   true,
				HTTPOnly: true,
				SameSite: "strict",
			})
			log.Println(time.Since(start).Milliseconds())
			return c.Send(guestsJSON)
		}
	}
	log.Println(time.Since(start).Milliseconds())
	return c.Status(fiber.StatusNotFound).
		JSON(fiber.Map{"errorMessage": "I don't know who you are"})
}
