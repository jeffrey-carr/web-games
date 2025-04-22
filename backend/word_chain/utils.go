package wordchain

import (
	"math/rand"
	"time"
)

// NewLobbyCode generates a new random lobby code
func NewLobbyCode() string {
	rand.New(rand.NewSource(time.Now().UnixNano()))

	// Create a byte slice to hold our 4 letters
	code := make([]byte, 4)

	// Generate 4 random uppercase letters (ASCII A-Z: 65-90)
	for i := range 4 {
		code[i] = byte(rand.Intn(26) + 65) // 65 is ASCII 'A', 26 letters in alphabet
	}

	return string(code)
}

// CalculateDiff calculates number of different characters between two words
// If the words are not equal length, returns -1
func CalculateDiff(wordOne string, wordTwo string) int {
	if len(wordOne) != len(wordTwo) {
		return -1
	}

	diff := 0
	for i := range len(wordOne) {
		if wordOne[i] != wordTwo[i] {
			diff++
		}
	}

	return diff
}
