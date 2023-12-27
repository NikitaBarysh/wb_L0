package handler

import (
	"github.com/go-chi/chi/v5"
)

type Router struct {
	handler *Handler
}

func NewRouter(handler *Handler) *Router {
	return &Router{handler: handler}
}

func (rt *Router) Register() *chi.Mux {
	r := chi.NewRouter()

	r.Get("/order/{orderUID}", rt.handler.getOrder)

	return r
}
