package middleware

import (
	"net/http"
	"web_games/utils"
)

// Cors represents cors middleware
type Cors struct {
	config utils.Config
}

// NewCors creates a new cors middleware
func NewCors(config utils.Config) Cors {
	return Cors{config: config}
}

// GetName gets the name of the middleware
func (c Cors) GetName() string {
	return "CORS"
}

// Apply applies the cors middleware
func (c Cors) Apply(w http.ResponseWriter, r *http.Request) (http.ResponseWriter, *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", c.config.FrontendDomain)
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	return w, r
}
