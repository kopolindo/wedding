package backend

import (
	"encoding/csv"
	"fmt"
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
		passphrase, _ := generatePassphrase(dictionary)
		hash, err := generateFromPassword(passphrase)
		if err != nil {
			log.Printf("error during argon2 password generation: %s\n", err.Error())
		}
		g := models.Guest{
			FirstName: r[0],
			LastName:  r[1],
			UUID:      uuid.New(),
			Secret:    hash,
			Confirmed: false,
			Notes:     []byte{},
		}
		fmt.Println(g.FirstName, g.LastName, passphrase)
		GUESTS = append(GUESTS, g)
	}
}
