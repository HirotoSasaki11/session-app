package router

import (
	"context"
	"log"
	"net/http"
	"session-sample/server/application"
	"session-sample/server/application/model"
	"session-sample/server/lib"

	"github.com/go-chi/chi"
)

type User struct {
	User *application.User
}

func (u *User) Create(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	user := new(model.User)
	err := lib.BodyToJson(r, user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	log.Println(user)
	err = u.User.Create(ctx, user)
	if err != nil {
		w.Write([]byte(err.Error()))
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (u *User) GetByID(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	id := chi.URLParam(r, "user_id")
	user, err := u.User.GetByID(ctx, id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = lib.Json(w, user)
	if err != nil {
		return
	}
	w.WriteHeader(http.StatusOK)
}
