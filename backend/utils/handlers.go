package utils

import (
	"binoku/entities"
	"net/http"
)

// NewHandler creates a new handler, applying default middlewares
func NewHandler(config entities.Config, w http.ResponseWriter, r *http.Request, handle func(w http.ResponseWriter, r *http.Request)) {
	// Allow frontend through CORs
	w.Header().Set("Access-Control-Allow-Origin", config.FrontendDomain)
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	handle(w, r)
}
