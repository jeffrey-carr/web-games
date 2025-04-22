package wordchain

import (
	"encoding/json"
	"net/http"
)

// Handler reprenents a word ladder handler
type Handler interface {
	NewGame(w http.ResponseWriter, r *http.Request)
	CreateLobby(w http.ResponseWriter, r *http.Request)
	ValidateAnswer(w http.ResponseWriter, r *http.Request)
}

type handler struct {
	controller Controller
}

// NewHandler creates a new handler
func NewHandler(controller Controller) Handler {
	return handler{
		controller: controller,
	}
}

func (h handler) NewGame(w http.ResponseWriter, r *http.Request) {
	game, err := h.controller.CreateGame(r.Context())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	marshalledGame, err := json.Marshal(game)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(marshalledGame)
}

func (h handler) CreateLobby(w http.ResponseWriter, r *http.Request) {
	// TODO
	w.WriteHeader(http.StatusNotImplemented)
}

func (h handler) ValidateAnswer(w http.ResponseWriter, r *http.Request) {
	var validateRequest ValidateAnswerRequest
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&validateRequest)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	correct, updatedGame, err := h.controller.ValidateGuess(validateRequest.Guess, validateRequest.GameState)
	if err != nil {
		sendResponse(w, http.StatusInternalServerError, "Error validating guess")
		return
	}

	sendResponse(w, http.StatusOK, ValidateAnswerResponse{Correct: correct, UpdatedGame: updatedGame})
}

func sendResponse(w http.ResponseWriter, status int, data any) {
	w.WriteHeader(status)
	err := json.NewEncoder(w).Encode(&data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
