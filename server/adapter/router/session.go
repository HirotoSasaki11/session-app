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
	} else if user == nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("not found user."))
	}
	lib.Json(w, user)
}
func (s *Session) Get(w http.ResponseWriter, r *http.Request) {
	// id := chi.URLParam(r, "id")
	ctx := context.Background()
	user, err := s.Session.Get(ctx, w, r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	lib.Json(w, user)
	// var data []byte
	// err = json.Unmarshal(data, user)
	// if err != nil {
	// 	w.WriteHeader(http.StatusInternalServerError)
	// }
	// w.Write(data)
}
