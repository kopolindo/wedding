package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"
	"wedding/database"
	"wedding/log"
	"wedding/models"
	"wedding/telegram"

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
			Notes:     g.Notes,
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

	// start telegram routine
	// marshal data to actual JSON struct ([]byte)
	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Errorf("[telegram] error during marshalling of data: %s", err.Error())
	}

	// set parse_mode to MarkdownV2 for JSON visualization
	telegram.TGMessage.ParseMode = "MarkdownV2"

	// cast the marshalled JSON ([]byte) to string
	prettyJsonDataString, err := prettyPrintJSON(string(jsonData))
	jsonDataString := telegram.PrepareMarkup(prettyJsonDataString)
	if err != nil {
		log.Errorf("[telegram] error during prettyfication of JSON data: %s", err.Error())
	}
	// prepare the message
	notificationMessage := fmt.Sprintf(
		"%s %s \\(`%s`\\) ha appena inviato correttamente i seguenti dati:\n```json\n%s\n```",
		firstName,
		lastName,
		uuid,
		jsonDataString,
	)

	log.Infof("[telegram] sending notification")

	// send message
	err = telegram.SendNotification(notificationMessage)
	if err != nil {
		// log error but do not return
		log.Errorf("[telegram] failed to send notification: %s", err.Error())
	}

	// return
	return c.Status(fiber.StatusOK).
		JSON(fiber.Map{"status": "ok"})
}

// Function to pretty print a JSON string
func prettyPrintJSON(jsonStr string) (string, error) {
	var prettyJSON bytes.Buffer
	err := json.Indent(&prettyJSON, []byte(jsonStr), "", "  ")
	if err != nil {
		return "", fmt.Errorf("failed to pretty print JSON: %v", err)
	}
	return prettyJSON.String(), nil
}
