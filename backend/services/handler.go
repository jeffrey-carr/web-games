package services

import (
	"net/http"
	"web_games/entities"
	"web_games/utils"
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
	config             utils.Config
	defaultMiddlewares []entities.Middleware
}

// NewHandler creates a new Handler
func NewHandler(config utils.Config, defaultMiddlewares []entities.Middleware) Handler {
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
	localMiddlewares := make([]entities.Middleware, len(middlewares))
	copy(localMiddlewares, middlewares)

	http.HandleFunc(
		slug,
		func(w http.ResponseWriter, r *http.Request) {
			if !h.confirmMethod(r, method) {
				w.WriteHeader(http.StatusMethodNotAllowed)
				return
			}

			w, r = h.applyMiddlewares(w, r, localMiddlewares...)

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
