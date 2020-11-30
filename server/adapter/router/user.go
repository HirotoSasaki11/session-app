package router

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"session-sample/server/application"
	"session-sample/server/application/model"
	"strconv"

	"github.com/go-chi/chi"
)

type User struct {
	User *application.User
}

func (u *User) Create(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	var user model.User
	length, err := strconv.Atoi(r.Header.Get("Content-Length"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	body := make([]byte, length)
	_, err = r.Body.Read(body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	err = json.Unmarshal(body, &user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	log.Println(user)
	err = u.User.Create(ctx, &user)
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
