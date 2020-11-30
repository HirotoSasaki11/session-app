package application

import (
	"context"
	"net/http"
	"session-sample/server/application/model"
	"session-sample/server/application/repository"
)

type Session struct {
	Session repository.Session
	User    repository.User
}

func (s *Session) Get(ctx context.Context, w http.ResponseWriter, r *http.Request, id string) (*model.User, error) {
	return s.Session.Get(ctx, w, r, id)
}

func (s *Session) Set(w http.ResponseWriter, r *http.Request, user *model.User) error {
	return s.Session.Set(w, r, user)
}
