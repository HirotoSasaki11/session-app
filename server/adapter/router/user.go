package router

import (
	"net/http"
	"session-sample/server/application"
)

type User struct {
	User application.User
}

func (u *User) Create(w http.ResponseWriter, r *http.Request) {

}
func (u *User) GetByID(w http.ResponseWriter, r *http.Request) {
}
