package repository

import (
	"context"
	"net/http"
	"session-sample/server/application/model"
)

type Session interface {
	Get(ctx context.Context, w http.ResponseWriter, r *http.Request) (*model.User, error)
	Set(ctx context.Context, w http.ResponseWriter, r *http.Request, user *model.User) error
}
