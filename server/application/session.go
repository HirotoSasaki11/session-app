package application

import (
	"session-sample/server/application/model"
	"session-sample/server/application/repository"
)

type Session struct {
	Session repository.Session
	User    repository.User
}

func (s *Session) Get(token string) error {
	return s.Get(token)
}

func (s *Session) Set(id string) (*model.User, error) {
	return s.Set(id)
}

func createTokenID() string {
	return ""
}
