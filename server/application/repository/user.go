package repository

import (
	"context"
	"session-sample/server/application/model"
)

type User interface {
	Create(ctx context.Context, user *model.User) error
	GetByID(ctx context.Context, user *model.User) (*model.User, error)
	Exists(ctx context.Context, user *model.User) (bool, error)
}
