package main

import (
	"fmt"
	"net/http"
	"os"
	"web_games/binoku"
	"web_games/entities"
	"web_games/middleware"
	"web_games/services"
	"web_games/utils"
	wordchain "web_games/word_chain"

	"github.com/pkg/errors"
)

func main() {
	envStr, ok := os.LookupEnv(utils.EnvironmentDeploymentVariable)
	if !ok {
		panic("could not find environment variable")
	}
	environment := utils.EnvironmentDeployment(envStr)
	if utils.IsZero(environment) {
		panic("could not cast deployment environment variable")
	}

	var configFile string
	if environment == utils.DeploymentProd {
		configFile = "config.prod.yaml"
	} else {
		configFile = "config.dev.yaml"
	}

	// Read config file
	config, err := utils.ReadConfig(configFile)
	if err != nil {
		panic(errors.Wrap(err, "error reading config file"))
	}

	encryptionKey, err := utils.GenerateGCMKey()
	if err != nil {
		panic(errors.Wrap(err, "error generating encryption key"))
	}
	encryptionService := services.NewEncryption(encryptionKey)

	wordChainDictionary, err := utils.ReadJSONFile[wordchain.Dictionary]("data/word_chain_dictionary.json")
	if err != nil {
		panic(err)
	}

	container := DependencyContainer{
		BinokuController: binoku.NewGameManager(),
		WordLadderController: wordchain.NewController(
			config.WordLadder.MaxServers,
			config.WordLadder.MaxPlayersPerServer,
			wordChainDictionary,
			encryptionService,
		),
	}

	handleService := services.NewHandler(
		config,
		[]entities.Middleware{
			middleware.NewCors(config),
		},
	)

	handleService.Handle(
		"/",
		http.MethodGet,
		func(w http.ResponseWriter, _ *http.Request) {
			w.Write([]byte("hello world"))
		},
	)

	// Binoku
	binokuHandler := binoku.NewHandler(container.BinokuController)
	handleService.Handle(
		"/binoku/new-game",
		http.MethodGet,
		binokuHandler.NewGame,
	)
	handleService.Handle(
		"/binoku/validate-game",
		http.MethodPost,
		binokuHandler.ValidateBoard,
	)

	// Word Ladder
	wordLadderHandler := wordchain.NewHandler(container.WordLadderController)
	handleService.Handle(
		"/word-chain/new-game",
		http.MethodGet,
		wordLadderHandler.NewGame,
	)
	handleService.Handle(
		"/word-chain/validate-answer",
		http.MethodPost,
		wordLadderHandler.ValidateAnswer,
	)

	port := config.Port
	listenAddr := fmt.Sprintf(":%d", port)
	fmt.Printf("Server listening on %d\n", port)
	if environment == utils.DeploymentProd {
		err = http.ListenAndServeTLS(listenAddr, config.FullCertPath, config.PrivateKeyPath, nil)
	} else {
		err = http.ListenAndServe(listenAddr, nil)
	}

	if err != nil {
		panic(err)
	}
}
