package main

import (
	"binoku/binoku"
	"binoku/entities"
	"binoku/middleware"
	"binoku/services"
	"binoku/utils"
	"fmt"
	"net/http"
	"os"

	"github.com/pkg/errors"
)

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
		configFile = "config.prod.yaml"
	} else {
		configFile = "config.dev.yaml"
	}

	// Read config file
	config, err := utils.ReadConfig(configFile)
	if err != nil {
		panic(errors.Wrap(err, "error reading config file"))
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
	binokuHandler := binoku.NewHandler()
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
