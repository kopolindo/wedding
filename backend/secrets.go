package backend

import (
	"bufio"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"log"
	"math/big"
	"os"
	"strconv"
	"strings"

	"github.com/trustelem/zxcvbn"
	"golang.org/x/crypto/argon2"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

const (
	FILENAME      string = "./passphrase-generator-dictionary.txt"
	MINSCORE      int    = 3
	PASSPHRASELEN int    = 2
)

type params struct {
	memory      uint32
	iterations  uint32
	parallelism uint8
	saltLength  uint32
	keyLength   uint32
}

var (
	argonParams = &params{
		memory:      64 * 1024,
		iterations:  3,
		parallelism: 2,
		saltLength:  16,
		keyLength:   32,
	}
)

// readDictionary reads the dictionary file and returns a slice of words
func readDictionary() ([]string, error) {
	file, err := os.Open(FILENAME)
	if err != nil {
		log.Printf("error during opening dictionary file: %s\n", err.Error())
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
		log.Printf("error during scanning of the dictionary file content: %s\n", err.Error())
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
			randomIndex := randomInt(int64(len(dictionary)))
			word := dictionary[randomIndex]
			word = cases.Title(language.Italian).String(word)
			for usedWords[word] {
				randomIndex = randomInt(int64(len(dictionary)))
				word = dictionary[randomIndex]
				word = cases.Title(language.Italian).String(word)
			}
			if i == PASSPHRASELEN-1 {
				randomNumber := randomInt(100)
				word += strconv.Itoa(int(randomNumber))
			}
			passphrase = append(passphrase, word)
		}
		result := zxcvbn.PasswordStrength(strings.Join(passphrase, "-"), nil)
		score = result.Score
	}
	return strings.Join(passphrase, "-"), score
}

func generateFromPassword(password string) (encodedHash string, err error) {
	// Generate a cryptographically secure random salt.
	salt, err := generateRandomBytes(argonParams.saltLength)
	if err != nil {
		return "", err
	}
	// Pass the plaintext password, salt and parameters to the argon2.IDKey
	// function. This will generate a hash of the password using the Argon2id
	// variant.
	hash := argon2.IDKey(
		[]byte(password),
		salt,
		argonParams.iterations,
		argonParams.memory,
		argonParams.parallelism,
		argonParams.keyLength,
	)

	// Base64 encode the salt and hashed password.
	b64Salt := base64.RawStdEncoding.EncodeToString(salt)
	b64Hash := base64.RawStdEncoding.EncodeToString(hash)

	// Return a string using the standard encoded hash representation.
	encodedHash = fmt.Sprintf(
		"$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s",
		argon2.Version,
		argonParams.memory,
		argonParams.iterations,
		argonParams.parallelism,
		b64Salt,
		b64Hash,
	)
	return encodedHash, nil
}

/*
	generateRandomBytes returns a byte slice

input: uint32, length of the slice
output: []byte, error
*/
func generateRandomBytes(n uint32) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}

	return b, nil
}

/*
	randomInt returns a random int64 value given a limsup

input: int64 the maximum value
output int64 the random value
*/
func randomInt(max int64) int64 {
	randomIndexBigInt, err := rand.Int(rand.Reader, big.NewInt(max))
	if err != nil {
		log.Printf("error during random int generation: %s\n", err.Error())
	}
	return randomIndexBigInt.Int64()
}
