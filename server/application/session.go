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

func (s *Session) Get(ctx context.Context, w http.ResponseWriter, r *http.Request) (*model.User, error) {
	return s.Session.Get(ctx, w, r)
}

func (s *Session) Set(ctx context.Context, w http.ResponseWriter, r *http.Request, id string) (*model.User, error) {
	user, err := s.User.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	err = s.Session.Set(ctx, w, r, user)
	if err != nil {
		return nil, err
	}
	return user, nil
}
