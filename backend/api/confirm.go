package api

import (
	"fmt"
	"strings"
	"time"
	"wedding/src/database"
	"wedding/src/log"
	"wedding/src/models"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func (v XValidator) Validate(data interface{}) []ErrorResponse {
	validationErrors := []ErrorResponse{}
	errs := validate.Struct(data)
	if errs != nil {
		for _, err := range errs.(validator.ValidationErrors) {
			// In this case data object is actually holding the User struct
			var elem ErrorResponse

			elem.FailedField = err.Field() // Export struct field name
			elem.Tag = err.Tag()           // Export struct tag
			elem.Value = err.Value()       // Export field value
			elem.Error = true

			validationErrors = append(validationErrors, elem)
		}
	}

	return validationErrors
}

// handleFormPost handles form submission
func handleFormPost(c *fiber.Ctx) error {
	validate = validator.New()
	myValidator := &XValidator{
		validator: validate,
	}

	sessionID := c.Cookies("session")
	if sessionID == "" {
		return c.Status(fiber.StatusInternalServerError).
			JSON(fiber.Map{"errorMessage": "Hai trovato il modo di superare il cane a tre teste! Ma non supererai me..."})
	}
	uuid, err := uuid.Parse(sessionID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).
			JSON(fiber.Map{"errorMessage": err.Error()})
	}

	data := new(JSONConfirmationForm)

	if err := c.BodyParser(data); err != nil {
		log.Errorf("error during JSON parsing: %s\n", err.Error())
		return c.Status(fiber.StatusInternalServerError).
			JSON(fiber.Map{"errorMessage": err.Error()})
	}

	var guests []models.Guest
	// Create guest struct from parsed data
	for _, g := range data.People {
		guest := models.Guest{
			ID:        g.ID,
			FirstName: g.FirstName,
			LastName:  g.LastName,
			UUID:      uuid,
			Confirmed: true,
			Notes:     []byte(g.Notes),
		}
		// Validation
		if errs := myValidator.Validate(&guest); len(errs) > 0 && errs[0].Error {
			errMsgs := make([]string, 0)

			for _, err := range errs {
				errMsgs = append(errMsgs, fmt.Sprintf(
					"[%s]: '%v' needs to implement '%s'",
					err.FailedField,
					err.Value,
					err.Tag,
				))
			}

			return &fiber.Error{
				Code:    fiber.ErrBadRequest.Code,
				Message: strings.Join(errMsgs, "\n"),
			}
		}
		guests = append(guests, guest)
	}

	for _, guest := range guests {
		if database.GuestExists(guest.ID, guest.UUID) {
			err = database.UpdateGuest(guest)
			if err != nil {
				log.Errorf("error after updating guest: %s\n", err.Error())
				return c.Status(fiber.StatusInternalServerError).
					JSON(fiber.Map{"errorMessage": err.Error()})
			}
			c.Cookie(&fiber.Cookie{
				Name:     "confirmed",
				Value:    "true",
				Expires:  time.Now().Add(24 * 7 * time.Hour),
				Secure:   true,
				HTTPOnly: false,
				SameSite: "strict",
			})
		} else {
			guest.ID = 0
			_, err := database.CreateGuest(guest)
			if err != nil {
				log.Errorf("error during guest creation: %s\n", err.Error())
				return c.Status(fiber.StatusInternalServerError).
					JSON(fiber.Map{"errorMessage": err.Error()})
			}
		}
	}

	return c.Status(fiber.StatusOK).
		JSON(fiber.Map{"status": "ok"})
}
