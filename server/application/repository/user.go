package repository

import (
	"session-sample/server/application/model"
)

type User interface {
	Create(user model.User) error
	GetByID(id string) (*model.User, error)
}
