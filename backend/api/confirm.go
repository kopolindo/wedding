package api

import (
	"encoding/json"
	"fmt"
	"time"
	"wedding/database"
	"wedding/log"
	"wedding/models"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type ValidationErrorJSON struct {
	Errors []struct {
		FailedField string `json:"failed_field"`
		Tag         string `json:"tag"`
		Value       string `json:"value"`
		Message     string `json:"message"`
	} `json:"errors"`
}

var FieldsMapping = map[string]string{
	"FirstName": "Nome",
	"LastName":  "Cognome",
	"Note":      "Notes",
}

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

	firstName, lastName, err := database.GetMainUserByUUID(uuid)
	if err != nil {
		log.Errorf("error during fetching of user details from UUID %s: %s", uuid, err.Error())
	}

	data := new(JSONConfirmationForm)

	if err := c.BodyParser(data); err != nil {
		log.Errorf("error during JSON parsing: %s", err.Error())
		return c.Status(fiber.StatusInternalServerError).
			JSON(fiber.Map{"errorMessage": err.Error()})
	}

	log.Debugf("received this data from %s %s (%s): %v", firstName, lastName, uuid, data)

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
			var validationErrorJSON ValidationErrorJSON

			for _, err := range errs {
				var message string
				switch err.Tag {
				case "min":
					message = fmt.Sprintf("Il campo %s deve avere una lunghezza minima di 3 caratteri", FieldsMapping[err.FailedField])
				case "max":
					message = fmt.Sprintf("Il campo %s deve avere una lunghezza massima di 20 caratteri", FieldsMapping[err.FailedField])
				case "ascii":
					message = fmt.Sprintf("Il campo %s può contenere solo caratteri ASCII validi", FieldsMapping[err.FailedField])
				case "required":
					message = fmt.Sprintf("Il campo %s è necessario", FieldsMapping[err.FailedField])
				default:
					message = fmt.Sprintf("Il campo %s presenta una non conformità imprevista: tag: %s, error: %s", FieldsMapping[err.FailedField], err.Tag, err.Value)
				}
				validationErrorJSON.Errors = append(validationErrorJSON.Errors, struct {
					FailedField string `json:"failed_field"`
					Tag         string `json:"tag"`
					Value       string `json:"value"`
					Message     string `json:"message"`
				}{
					FailedField: err.FailedField,
					Tag:         err.Tag,
					Value:       fmt.Sprintf("%s", err.Value),
					Message:     message,
				})
			}

			errorJSON, err := json.Marshal(validationErrorJSON)
			if err != nil {
				log.Errorf("error marshaling validation errors: %v", err)
				return &fiber.Error{
					Code:    fiber.StatusInternalServerError,
					Message: "Internal Server Error",
				}
			}
			log.Errorf(string(errorJSON))
			c.Set("Content-Type", "application/json")
			return &fiber.Error{
				Code:    fiber.ErrBadRequest.Code,
				Message: string(errorJSON),
			}
		}
		guests = append(guests, guest)
	}

	for _, guest := range guests {
		if database.GuestExists(guest.ID, guest.UUID) {
			err = database.UpdateGuest(guest)
			if err != nil {
				log.Errorf("error after updating guest: %s", err.Error())
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
				log.Errorf("error during guest creation: %s", err.Error())
				return c.Status(fiber.StatusInternalServerError).
					JSON(fiber.Map{"errorMessage": err.Error()})
			}
		}
	}

	log.Infof("POST request correctly handled for user: %s, %s, %s", uuid, firstName, lastName)
	return c.Status(fiber.StatusOK).
		JSON(fiber.Map{"status": "ok"})
}
