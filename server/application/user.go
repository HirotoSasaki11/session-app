package application

import (
	"context"
	"log"
	"session-sample/server/application/model"
	"session-sample/server/application/repository"
)

type User struct {
	User repository.User
}

func (u *User) Create(ctx context.Context, user *model.User) error {
	exists, err := u.User.Exists(ctx, user)
	if err != nil {
		return err
	} else if exists {
		log.Println("bad request.")
		return nil
	}

	return u.User.Create(ctx, user)
}

func (u *User) GetByID(ctx context.Context, id string) (*model.User, error) {
	return u.GetByID(ctx, id)
}
