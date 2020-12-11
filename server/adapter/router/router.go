package router

import (
	"session-sample/server/registry"

	"github.com/go-chi/chi"
)

// type Router interface{}
type Provide struct {
	User    *User
	Session *Session
	Article *Article
}

func NewRouter(r registry.Resource, p Provide) *chi.Mux {
	mux := chi.NewRouter()
	mux.Route("/session", func(r chi.Router) {
		r.Get("/", p.Session.Get)
		r.Post("/{id}", p.Session.Set)
	})
	mux.Route("/users", func(r chi.Router) {
		r.Get("/{user_id}", p.User.GetByID)
		r.Post("/", p.User.Create)
	})
	mux.Route("/article", func(r chi.Router) {
		r.Get("/{article_id}", p.Article.GetByID)
		r.Post("/", p.Article.Create)
	})
	return mux
}
