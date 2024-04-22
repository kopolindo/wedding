package backend

import (
	"encoding/csv"
	"log"
	"os"
	"wedding/models"

	"github.com/google/uuid"
)

var GUESTS []models.Guest

const (
	guestsfile string = "guests.csv"
)

// createGuestList read the list of guest names (only per group) and stores it in GUESTS
func createGuestList() {
	// open file
	f, err := os.Open(guestsfile)
	if err != nil {
		log.Fatalf("error during file opening: %s\n", err.Error())
	}
	// create csv reader
	r := csv.NewReader(f)
	records, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	dictionary, err := readDictionary()
	if err != nil {
		log.Fatal(err.Error())
	}
	for _, r := range records {
		g := models.Guest{
			FirstName: r[0],
			LastName:  r[1],
			UUID:      uuid.New(),
			Confirmed: false,
			Notes:     []byte{},
		}
		GUESTS = append(GUESTS, g)
		passphrase, entropy := GeneratePassphrase(dictionary)
		log.Printf("%s (%d)\n", passphrase, entropy)
	}
}
