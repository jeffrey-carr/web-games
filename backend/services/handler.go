package services

import (
	"binoku/entities"
	"net/http"
)

// Handler represents a handler
type Handler interface {
	Handle(
		slug string,
		method string,
		handle func(w http.ResponseWriter, r *http.Request),
		middlewares ...entities.Middleware,
	)
}

type handler struct {
	config             entities.Config
	defaultMiddlewares []entities.Middleware
}

// NewHandler creates a new Handler
func NewHandler(config entities.Config, defaultMiddlewares []entities.Middleware) Handler {
	return handler{
		config:             config,
		defaultMiddlewares: defaultMiddlewares,
	}
}

func (h handler) Handle(
	slug string,
	method string,
	handle func(w http.ResponseWriter, r *http.Request),
	middlewares ...entities.Middleware,
) {
	http.HandleFunc(
		slug,
		func(w http.ResponseWriter, r *http.Request) {
			if !h.confirmMethod(r, method) {
				w.WriteHeader(http.StatusMethodNotAllowed)
				return
			}

			w, r = h.applyMiddlewares(w, r, middlewares...)

			handle(w, r)
		},
	)
}

func (h handler) confirmMethod(request *http.Request, method string) bool {
	return request.Method == method
}

func (h handler) applyMiddlewares(
	w http.ResponseWriter,
	r *http.Request,
	middlewares ...entities.Middleware,
) (http.ResponseWriter, *http.Request) {
	allMiddlewares := append(h.defaultMiddlewares, middlewares...)
	for _, middleware := range allMiddlewares {
		w, r = middleware.Apply(w, r)
	}

	return w, r
}
