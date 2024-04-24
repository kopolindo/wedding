package api

import (
	"fmt"
	"net/http"
	"net/url"
	"sync"
	"time"
	"wedding/backend"
	"wedding/database"
	"wedding/models"

	"github.com/gofiber/fiber/v2"
)

// handleSecret renders the login page
// method POST
// route /secret
func handleSecret(c *fiber.Ctx) error {
	start := time.Now()
	c.Type("html")
	// Parse form data
	secret := c.FormValue("secret")
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
			return c.JSON(fiber.Map{
				"errorMessage": res.err.Error(),
				"statusCode":   http.StatusInternalServerError,
			})
		}
		if res.ok {
			location, err := url.JoinPath("/guest", res.guest.UUID.String())
			if err != nil {
				return c.JSON(fiber.Map{
					"errorMessage": err.Error(),
					"statusCode":   http.StatusInternalServerError,
				})
			}
			fmt.Println(time.Since(start).Milliseconds())
			return c.Redirect(location)
		}
	}
	fmt.Println(time.Since(start).Milliseconds())
	return c.Redirect("/chisono")
}
