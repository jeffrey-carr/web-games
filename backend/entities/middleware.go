package entities

import "net/http"

type Middleware interface {
	Apply(w http.ResponseWriter, r *http.Request) (http.ResponseWriter, *http.Request)
}
