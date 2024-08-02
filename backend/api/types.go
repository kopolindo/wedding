package api

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var (
	Router   http.Handler
	App      *fiber.App
	validate *validator.Validate
)

const (
	COOKIEPASSWORDFILE       = "./cookie-passphrase.txt"
	COOKIEPASSWORDFILEDOCKER = "/run/secrets/cookie_passphrase"
	SCHEMA                   = "http"
	DOMAIN                   = "localhost"
)

type (
	JSONConfirmationForm struct {
		Guests int `json:"guests"`
		People []struct {
			ID        uint   `json:"ID"`
			FirstName string `json:"first_name"`
			LastName  string `json:"last_name"`
			Notes     string `json:"notes"`
			Type      int    `json:"type"`
		} `json:"people"`
	}

	ErrorResponse struct {
		Error       bool
		FailedField string
		Tag         string
		Value       interface{}
	}

	XValidator struct {
		validator *validator.Validate
	}

	GlobalErrorHandlerResp struct {
		Success bool   `json:"success"`
		Message string `json:"message"`
	}

	secretRequestPayload struct {
		Secret string `json:"secret"`
	}

	secretResponseBody struct {
		UUID string `json:"uuid"`
	}

	GuestToDelete struct {
		ID uint `json:"id"`
	}

	// Used for security reasons. This struct does not hold the secret ;)
	responseGuest struct {
		ID        uint   `json:"id"`
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Confirmed bool   `json:"confirmed"`
		Notes     string `json:"notes"`
		Type      int    `json:"type"`
	}

	responseGuests struct {
		Guests []responseGuest `json:"guests"`
	}
)
