package api

import (
	"encoding/json"
	"sync"
	"time"
	"wedding/argon"
	"wedding/database"
	"wedding/models"

	"github.com/gofiber/fiber/v2"
)

// handleSecret renders the login page
// method POST
// route /secret
func handleSecret(c *fiber.Ctx) error {
	// Parse form data
	payload := &secretRequestPayload{}
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
			if guest.Secret != "" {
				ok, err := argon.ComparePasswordAndHash(secret, guest.Secret)
				results <- result{guest, ok, err}
			}
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
			response := secretResponseBody{
				UUID: uuid.String(),
			}
			responseJSON, err := json.Marshal(response)
			if err != nil {
				return c.Status(fiber.StatusInternalServerError).
					JSON(fiber.Map{"errorMessage": err.Error()})
			}
			c.Cookie(&fiber.Cookie{
				Name:     "session",
				Value:    uuid.String(),
				Expires:  time.Now().Add(24 * 7 * time.Hour),
				Secure:   true,
				HTTPOnly: true,
				SameSite: "strict",
			})
			// this second cookie is not really important for authentication purpose
			// it's only required for client-side rendering of "private" pages
			// in React nothing is "really" private.
			// HTTPOnly is false because client-side code needs to check the presence
			// of auth cookie to render pages. Even if users forge such cookie, it is
			// not used server-side to authenticate requests.
			c.Cookie(&fiber.Cookie{
				Name:     "auth",
				Value:    "true",
				Expires:  time.Now().Add(24 * 7 * time.Hour),
				Secure:   true,
				HTTPOnly: false,
				SameSite: "strict",
			})
			cookie, err := confirmedCookie(uuid)
			if err != nil {
				return c.Status(fiber.StatusInternalServerError).
					JSON(fiber.Map{"errorMessage": err.Error()})
			}
			c.Cookie(&cookie)
			c.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)
			return c.Send(responseJSON)
		}
	}
	return c.Status(fiber.StatusNotFound).
		JSON(fiber.Map{"errorMessage": "Parola d'ordine sbagliata. Neville vai a chiamare Hermione e riprova"})
}
