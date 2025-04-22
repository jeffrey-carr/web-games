package entities

import "net/http"

type Middleware interface {
	// For debugging
	GetName() string
	Apply(w http.ResponseWriter, r *http.Request) (http.ResponseWriter, *http.Request)
}
