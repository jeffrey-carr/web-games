package main

import (
	"binoku/controllers"
	"binoku/entities"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func handler(w http.ResponseWriter, r *http.Request, handle func(w http.ResponseWriter, r *http.Request)) {
	// Allow frontend through CORs
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	handle(w, r)
}

func validateSize(size int) (bool, string) {
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

func handleNewGame(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

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

	isSizeValid, validationMessage := validateSize(size)
	if !isSizeValid {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(validationMessage))
		return
	}

	controller := controllers.NewGameManager()
	board, err := controller.GenerateBoard(size)
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

func handleValidateAnswer(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	if r.Body == nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Get request
	var validateRequest entities.ValidateGameRequest
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&validateRequest)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	controller := controllers.NewGameManager()

	// Validate board
	isValid := controller.ValidateBoard(validateRequest.Board)

	response := entities.ValidateGameResponse{
		Valid: isValid,
	}
	marshalledResponse, err := json.Marshal(response)
	if err != nil {
		fmt.Printf("Error marshalling response: %s\n", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(marshalledResponse)
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		handler(w, r, func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("hello world"))
		})
	})
	http.HandleFunc("/new-game", func(w http.ResponseWriter, r *http.Request) { handler(w, r, handleNewGame) })
	http.HandleFunc("/validate-game", func(w http.ResponseWriter, r *http.Request) { handler(w, r, handleValidateAnswer) })

	port := 3001
	fmt.Printf("Server listening on %d\n", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
