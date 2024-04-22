package backend

import (
	"bufio"
	"log"
	"math"
	"math/rand"
	"os"
	"strconv"
	"strings"

	"github.com/trustelem/zxcvbn"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

const (
	FILENAME      string = "./passphrase-generator-dictionary.txt"
	MINSCORE      int    = 3
	PASSPHRASELEN int    = 2
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

// GeneratePassphrase creates a passphrase of specified length from the dictionary
func GeneratePassphrase(dictionary []string) (string, int) {
	var passphrase []string
	score := 0
	for score < MINSCORE {
		passphrase = nil
		usedWords := make(map[string]bool)
		for i := 0; i < PASSPHRASELEN; i++ {
			randomIndex := rand.Intn(len(dictionary))
			word := dictionary[randomIndex]
			word = cases.Title(language.Italian).String(word)
			for usedWords[word] {
				randomIndex = rand.Intn(len(dictionary))
				word = dictionary[randomIndex]
				word = cases.Title(language.Italian).String(word)
			}
			if i == PASSPHRASELEN-1 {
				randomNumber := rand.Intn(100)
				word += strconv.Itoa(randomNumber)
			}
			passphrase = append(passphrase, word)
		}
		result := zxcvbn.PasswordStrength(strings.Join(passphrase, "-"), nil)
		score = result.Score
	}
	return strings.Join(passphrase, "-"), score
}

func Shannon(value string) (bits int) {
	frq := make(map[rune]float64)

	//get frequency of characters
	for _, i := range value {
		frq[i]++
	}

	var sum float64

	for _, v := range frq {
		f := v / float64(len(value))
		sum += f * math.Log2(f)
	}

	bits = int(math.Ceil(sum*-1)) * len(value)
	return
}
