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

type response struct {
	UUID string `json:"uuid"`
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
			response := response{
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
			log.Println(time.Since(start).Milliseconds())
			c.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)
			return c.Send(responseJSON)
		}
	}
	log.Println(time.Since(start).Milliseconds())
	return c.Status(fiber.StatusNotFound).
		JSON(fiber.Map{"errorMessage": "Parola d'ordine sbagliata. Neville vai a chiamare Hermione e riprova"})
}
