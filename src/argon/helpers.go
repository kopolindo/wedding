package argon

import (
	"crypto/rand"
	"log"
	"math/big"
)

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
func RandomInt(max int64) int64 {
	randomIndexBigInt, err := rand.Int(rand.Reader, big.NewInt(max))
	if err != nil {
		log.Printf("error during random int generation: %s\n", err.Error())
	}
	return randomIndexBigInt.Int64()
}
