package entities

// GamePiece represents a valid game piece
type GamePiece int

const (
	// Zero represents a game piece of 0
	Zero GamePiece = 0
	// One represents a game piece of 1
	One GamePiece = 1
	// Empty represents an empty space on the game board
	Empty GamePiece = -1

	// EasyDifficulty is the percentage of empty spaces at the start of a game
	EasyDifficulty = 0.45
	// MediumDifficulty is the percentage of empty spaces at the start of a game
	MediumDifficulty = 0.95
	// HardDifficulty is the percentage of empty spaces at the start of a game
	HardDifficulty = 0.25
)

// Game represents a game object
type Game struct {
	Board [][]int `json:"board"`
}

// Coordinate represents a space on the board
type Coordinate struct {
	Col int `json:"x"`
	Row int `json:"y"`
}
