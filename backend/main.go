package main

import (
	"binoku/controllers"
	"binoku/entities"
	"binoku/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/pkg/errors"
)

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
	envStr, ok := os.LookupEnv(entities.EnvironmentDeploymentVariable)
	if !ok {
		panic("could not find environment variable")
	}
	environment := entities.EnvironmentDeployment(envStr)
	if utils.IsZero(environment) {
		panic("could not cast deployment environment variable")
	}

	var configFile string
	if environment == entities.DeploymentProd {
		fmt.Println("Using production config")
		configFile = "config.prod.yaml"
	} else {
		configFile = "config.dev.yaml"
	}

	// Read config file
	config, err := utils.ReadConfig(configFile)
	if err != nil {
		panic(errors.Wrap(err, "error reading config file"))
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		utils.NewHandler(config, w, r, func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("hello world"))
		})
	})
	http.HandleFunc("/new-game", func(w http.ResponseWriter, r *http.Request) { utils.NewHandler(config, w, r, handleNewGame) })
	http.HandleFunc("/validate-game", func(w http.ResponseWriter, r *http.Request) { utils.NewHandler(config, w, r, handleValidateAnswer) })

	port := config.Port
	listenAddr := fmt.Sprintf(":%d", port)
	fmt.Printf("Server listening on %d\n", port)
	if environment == entities.DeploymentProd {
		err = http.ListenAndServeTLS(listenAddr, config.FullCertPath, config.PrivateKeyPath, nil)
	} else {
		err = http.ListenAndServe(listenAddr, nil)
	}

	if err != nil {
		panic(err)
	}
}
