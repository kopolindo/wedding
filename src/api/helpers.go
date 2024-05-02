package api

import (
	"log"
	"os"
)

const COOKIEPASSWORDFILE = "../cookie-passphrase.txt"

// readCookiePassword initializes COOKIEPASSWORD variable with content of COOKIEPASSWORDFILE
func readCookiePassword() {
	content, err := os.ReadFile(COOKIEPASSWORDFILE)
	if err != nil {
		log.Fatalf("Error reading password file. %s\n", err.Error())
		return
	}
	COOKIEPASSWORD = string(content)
}
