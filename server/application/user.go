package application

import (
	"session-sample/server/application/model"
)

type UserInterface interface {
	Create(user model.User) error
	GetByID(id string) (*model.User, error)
}

type User struct {
	Interface UserInterface
}

func (u *User) Create(user model.User) error {
	return u.Interface.Create(user)
}

func (u *User) GetByID(id string) (*model.User, error) {
	return u.Interface.GetByID(id)
}
