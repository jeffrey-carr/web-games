package binoku

import (
	"encoding/json"
	"net/http"
	"strconv"
)

// Handler represents a Binoku handler
type Handler interface {
	NewGame(w http.ResponseWriter, r *http.Request)
	ValidateBoard(w http.ResponseWriter, r *http.Request)
}

type handler struct {
	controller GameManager
}

// NewHandler creates a new Binoku handler
func NewHandler(controller GameManager) Handler {
	return handler{
		controller: controller,
	}
}

// NewGame handles a new game request
func (h handler) NewGame(w http.ResponseWriter, r *http.Request) {
	// Get board sizeParam
	sizeParam := r.URL.Query().Get("size")
	if sizeParam == "" {
		// Default board size
		sizeParam = "6"
	}

	size, err := strconv.Atoi(sizeParam)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	isSizeValid, validationMessage := h.validateSize(size)
	if !isSizeValid {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(validationMessage))
		return
	}

	board, err := h.controller.GenerateBoard(size)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	marshalledBoard, err := json.Marshal(board)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(marshalledBoard)
}

// ValidateBoard validates a user's board for correctness
func (h handler) ValidateBoard(w http.ResponseWriter, r *http.Request) {
	if r.Body == nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Get request
	var validateRequest ValidateGameRequest
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&validateRequest)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Validate board
	isValid, invalidHint := h.controller.ValidateBoard(validateRequest.Board)

	response := ValidateGameResponse{
		Valid: isValid,
		Hint:  invalidHint,
	}
	marshalledResponse, err := json.Marshal(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(marshalledResponse)
}

func (h handler) validateSize(size int) (bool, string) {
	if size%2 != 0 {
		return false, "Board size must be even"
	}
	if size < 4 {
		return false, "Minimum board size is 4"
	}
	if size > 8 {
		return false, "Maximum board size is 8"
	}

	return true, ""
}
