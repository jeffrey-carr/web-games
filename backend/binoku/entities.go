package binoku

// GamePiece represents a valid game piece
type GamePiece int

const (
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
	Board [][]GamePiece `json:"board"`
}

// Coordinate represents a space on the board
type Coordinate struct {
	Col int `json:"x"`
	Row int `json:"y"`
}

// InvalidBoardHint is the hint shown to the user when their board is invalid
type InvalidBoardHint struct {
	Rows []int `json:"rows"`
	Cols []int `json:"cols"`
}

// ValidateGameRequest is the requests to validate a completed board
type ValidateGameRequest struct {
	Board [][]GamePiece `json:"board"`
}

// ValidateGameResponse is the resonse to validate a completed board
type ValidateGameResponse struct {
	Valid bool             `json:"valid"`
	Hint  InvalidBoardHint `json:"hint,omitempty"`
}
