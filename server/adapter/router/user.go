package router

import (
	"context"
	"encoding/json"
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
	lib.BodyToJson(r, user)
	log.Println(user)
	err := u.User.Create(ctx, user)
	if err != nil {
		w.Write([]byte(err.Error()))
		w.WriteHeader(http.StatusBadRequest)
	}
	w.WriteHeader(http.StatusCreated)
}
func (u *User) GetByID(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	id := chi.URLParam(r, "user_id")
	user, err := u.User.GetByID(ctx, id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	result, err := json.Marshal(user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	w.Write(result)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	w.WriteHeader(http.StatusOK)
}
