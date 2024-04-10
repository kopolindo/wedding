package database

import (
	"log"
	"os"
)

func readUserPassword() {
	content, err := os.ReadFile(USERPASSWORDFILE)
	if err != nil {
		log.Fatalf("Error reading password file. %s\n", err.Error())
		return
	}
	USERPASSWORD = string(content)
}
