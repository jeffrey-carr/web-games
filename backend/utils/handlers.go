package utils

import (
	"binoku/entities"
	"net/http"
)

// NewHandler creates a new handler, applying default middlewares
func NewHandler(config entities.Config, w http.ResponseWriter, r *http.Request, handle func(w http.ResponseWriter, r *http.Request)) {
	// Allow frontend through CORs
	origin := "http://localhost"
	if config.Environment == entities.DeploymentProd {
		origin = config.FrontendDomain
	}
	w.Header().Set("Access-Control-Allow-Origin", origin)
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	handle(w, r)
}
