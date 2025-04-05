package middleware

import (
	"binoku/entities"
	"net/http"
)

// Cors represents cors middleware
type Cors struct {
	config entities.Config
}

// NewCors creates a new cors middleware
func NewCors(config entities.Config) Cors {
	return Cors{config: config}
}

// Apply applies the cors middleware
func (c Cors) Apply(w http.ResponseWriter, r *http.Request) (http.ResponseWriter, *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", c.config.FrontendDomain)
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	return w, r
}
