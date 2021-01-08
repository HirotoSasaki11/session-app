package router

import (
	"context"
	"net/http"
	"session-sample/server/application"
	"session-sample/server/lib"

	"github.com/go-chi/chi"
)

type Session struct {
	Session *application.Session
}

func (s *Session) Set(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	ctx := context.Background()
	user, err := s.Session.Set(ctx, w, r, id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	} else if user == nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("not found user."))
		return
	}
	lib.Json(w, user)
}
func (s *Session) Get(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	user, err := s.Session.Get(ctx, w, r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	} else if user != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	w.WriteHeader(http.StatusOK)
	// lib.Json(w, user)
}

func (s *Session) Security(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.Background()
		user, err := s.Session.Get(ctx, w, r)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		} else if user == nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		next.ServeHTTP(w, r)
	})
}
