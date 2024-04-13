package database

import (
	"log"
	"os"
)

// readUserPassword initializes USERPASSWORD variable with content of USERPASSWORDFILE
func readUserPassword() {
	content, err := os.ReadFile(USERPASSWORDFILE)
	if err != nil {
		log.Fatalf("Error reading password file. %s\n", err.Error())
		return
	}
	USERPASSWORD = string(content)
}
