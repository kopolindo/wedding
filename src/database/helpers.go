package database

import (
	"bufio"
	"encoding/csv"
	"os"
	"strconv"
	"strings"
	"wedding/src/argon"
	"wedding/src/log"
	"wedding/src/models"

	"github.com/google/uuid"
	"github.com/trustelem/zxcvbn"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

var (
	GUESTS []models.Guest
)

const (
	guestsfile    string = "../guests.csv"
	FILENAME      string = "../passphrase-generator-dictionary.txt"
	MINSCORE      int    = 3
	PASSPHRASELEN int    = 2
)

// readDictionary reads the dictionary file and returns a slice of words
func readDictionary() ([]string, error) {
	file, err := os.Open(FILENAME)
	if err != nil {
		log.Errorf("error during opening dictionary file: %s\n", err.Error())
		return nil, err
	}
	defer file.Close()

	var dictionary []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		word := strings.TrimSpace(scanner.Text())
		if word != "" {
			dictionary = append(dictionary, word)
		}
	}
	if err := scanner.Err(); err != nil {
		log.Errorf("error during scanning of the dictionary file content: %s\n", err.Error())
		return nil, err
	}
	return dictionary, nil
}

// generatePassphrase creates a passphrase of specified length from the dictionary
func generatePassphrase(dictionary []string) (string, int) {
	var passphrase []string
	score := 0
	for score < MINSCORE {
		passphrase = nil
		usedWords := make(map[string]bool)
		for i := 0; i < PASSPHRASELEN; i++ {
			randomIndex := argon.RandomInt(int64(len(dictionary)))
			word := dictionary[randomIndex]
			word = cases.Title(language.Italian).String(word)
			for usedWords[word] {
				randomIndex = argon.RandomInt(int64(len(dictionary)))
				word = dictionary[randomIndex]
				word = cases.Title(language.Italian).String(word)
			}
			if i == PASSPHRASELEN-1 {
				randomNumber := argon.RandomInt(100)
				word += strconv.Itoa(int(randomNumber))
			}
			passphrase = append(passphrase, word)
		}
		result := zxcvbn.PasswordStrength(strings.Join(passphrase, "-"), nil)
		score = result.Score
	}
	return strings.Join(passphrase, "-"), score
}

// readUserPassword initializes USERPASSWORD variable with content of USERPASSWORDFILE
func readUserPassword() {
	content, err := os.ReadFile(USERPASSWORDFILE)
	if err != nil {
		log.Errorf("Error reading password file. %s\n", err.Error())
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
		log.Errorf(
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

// createGuestList read the list of guest names (only per group) and stores it in GUESTS
func createGuestList() {
	var guests Guests
	// open file
	f, err := os.Open(guestsfile)
	if err != nil {
		log.Errorf("error during file opening: %s\n", err.Error())
	}
	// create csv reader
	r := csv.NewReader(f)
	records, err := r.ReadAll()
	if err != nil {
		log.Errorf(err.Error())
	}
	dictionary, err := readDictionary()
	if err != nil {
		log.Errorf(err.Error())
	}
	for _, r := range records {
		passphrase, _ := generatePassphrase(dictionary)
		hash, err := argon.GenerateFromPassword(passphrase)
		if err != nil {
			log.Errorf("error during argon2 password generation: %s\n", err.Error())
		}
		u := uuid.New()
		g := models.Guest{
			FirstName: r[0],
			LastName:  r[1],
			UUID:      u,
			Secret:    hash,
			Confirmed: false,
			Notes:     []byte{},
		}
		guest := &Guest{
			FirstName:  g.FirstName,
			LastName:   g.LastName,
			Passphrase: passphrase,
			UUID:       u.String(),
		}
		guests = append(guests, *guest)
		GUESTS = append(GUESTS, g)
	}
	writeToCsv(&guests)
}

// Define a struct to hold your data
type Guest struct {
	FirstName  string
	LastName   string
	Passphrase string
	UUID       string
}

type Guests []Guest

func writeToCsv(guests *Guests) {
	file, err := os.Create("created_guest.csv")
	if err != nil {
		log.Errorf("CSV file creation error:", err.Error())
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	err = writer.Write([]string{"First Name", "Last Name", "UUID", "Passphrase"})
	if err != nil {
		panic(err)
	}

	for _, guest := range *guests {
		err = writer.Write([]string{guest.FirstName, guest.LastName, guest.UUID, guest.Passphrase})
		if err != nil {
			log.Errorf("writing error:", err.Error())
		}
	}

}
