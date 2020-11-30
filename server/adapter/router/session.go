package router

import (
	"context"
	"encoding/json"
	"net/http"
	"session-sample/server/application"

	"github.com/go-chi/chi"
)

type Session struct {
	Session *application.Session
}

func (s *Session) Set(w http.ResponseWriter, r *http.Request) {

}
func (s *Session) Get(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	ctx := context.Background()
	user, err := s.Session.Get(ctx, w, r, id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	var data []byte
	err = json.Unmarshal(data, user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.Write(data)

}
