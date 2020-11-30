package repository

import (
	"context"
	"net/http"
	"session-sample/server/application/model"
)

type Session interface {
	Get(ctx context.Context, w http.ResponseWriter, r *http.Request, id string) (*model.User, error)
	Set(w http.ResponseWriter, r *http.Request, user *model.User) error
}
