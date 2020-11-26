package router

import (
	"net/http"
	"session-sample/server/application"
)

type Session struct {
	Session *application.Session
}

func (s *Session) Set(w http.ResponseWriter, r *http.Request) {

}
func (s *Session) Get(w http.ResponseWriter, r *http.Request) {

}
