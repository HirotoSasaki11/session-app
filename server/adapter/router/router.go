package router

import (
	"session-sample/server/registry"

	"github.com/go-chi/chi"
)

// type Router interface{}
type Provide struct {
	User    *User
	Session *Session
}

func NewRouter(r registry.Resource, p Provide) *chi.Mux {
	mux := chi.NewRouter()
	mux.Post("/session", p.Session.Get)
	mux.Route("/users", func(r chi.Router) {
		r.Get("/{userid}", p.User.GetByID)
		r.Post("/", p.User.Create)
	})
	return mux
}
