package database

import (
	"log"
	"os"
	"wedding/src/models"
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

/*
isTableEmpty returns true if table is empty, false otherwise
input: string (table name)
*/
func isTableEmpty(table string) bool {
	isEmpty := false
	guests := &[]models.Guest{}
	result := db.Table(table).Find(guests)
	if result.Error != nil {
		log.Printf(
			"error while checking if table %s is empty (%s)\n",
			table,
			result.Error.Error(),
		)
	}
	if result.RowsAffected == int64(len(*guests)) && len(*guests) == 0 {
		isEmpty = true
	}
	return isEmpty
}
