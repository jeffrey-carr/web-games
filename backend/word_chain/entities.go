package wordchain

// Dictionary is the type for the dictionary of the WordChain game
type Dictionary map[string][]string

// Chain is the type for the actual word chain
type Chain []string

// Game represents a word ladder game
type Game struct {
	GameState

	UUID string `json:"uuid"`
	// EncryptedState is the encrypted version of the game for
	// progress verification
	EncryptedState string `json:"encryptedState"`
}

// GameState is the relevant game state for Word Chain
type GameState struct {
	// GeneratedChain is the readable format of the game
	GeneratedChain Chain `json:"generatedChain"`
	// UserProgress is the word in the chain that the user is
	// currently guessing
	UserProgress int `json:"userProgress"`
}

// ValidateAnswerRequest is the body of an HTTP request to validate a guess
type ValidateAnswerRequest struct {
	Guess     string `json:"guess"`
	GameState Game   `json:"gameState"`
}

// ValidateAnswerResponse is the response give to a ValidateAnswerRequest
type ValidateAnswerResponse struct {
	Correct     bool `json:"correct"`
	UpdatedGame Game `json:"updatedGame"`
}
