package entities

// ValidateGameRequest is the requests to validate a completed board
type ValidateGameRequest struct {
	Board [][]int `json:"board"`
}

// ValidateGameResponse is the resonse to validate a completed board
type ValidateGameResponse struct {
	Valid bool             `json:"valid"`
	Hint  InvalidBoardHint `json:"hint,omitempty"`
}
