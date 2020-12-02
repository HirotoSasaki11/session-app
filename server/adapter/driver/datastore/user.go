package datastore

import (
	"context"
	"log"
	"session-sample/server/application/model"
	"session-sample/server/config"

	"cloud.google.com/go/datastore"
)

type User struct {
	client *datastore.Client
	kind   string
}

func ProvideUser(client *datastore.Client) *User {
	return &User{client: client, kind: model.UsersEntityName}
}

func (u *User) NewKey(kind, name string) *datastore.Key {
	return &datastore.Key{
		Name:      name,
		Kind:      kind,
		Namespace: config.NameSpace,
	}
}
func (u *User) Create(ctx context.Context, user *model.User) error {
	_, err := u.client.Put(ctx, u.NewKey(u.kind, user.ID), user)
	if err != nil {
		return err
	}
	return nil
}

func (u *User) GetByID(ctx context.Context, id string) (*model.User, error) {
	user := new(model.User)
	err := u.client.Get(ctx, u.NewKey(u.kind, id), user)
	if err != nil {
		return nil, err
	}
	log.Println(user)
	return user, nil
}

func (u *User) GetByName(ctx context.Context, user *model.User) (*model.User, error) {
	var data []model.User
	q := datastore.NewQuery(u.kind).Namespace(config.NameSpace).Filter("Email", user.Email).Limit(1)
	_, err := u.client.GetAll(ctx, q, &data)
	if err != nil {
		return nil, err
	}
	return &data[0], nil
}

func (u *User) Exists(ctx context.Context, user *model.User) (bool, error) {
	existsID, err := u.ExistsID(ctx, user)
	if err != nil {
		return false, err
	}
	existsEmail, err := u.ExistsEmail(ctx, user)
	if err != nil {
		return false, err
	}
	if existsEmail || existsID {
		return true, nil
	}
	return false, nil
}
func (u *User) ExistsID(ctx context.Context, user *model.User) (bool, error) {
	data := new(model.User)
	key := u.NewKey(u.kind, user.ID)
	log.Println(key)
	err := u.client.Get(ctx, key, data)
	if err != nil {
		if err == datastore.ErrNoSuchEntity {
			return false, nil
		}
		return false, err
	} else if data != nil {
		log.Println("ID is exists")
		return true, nil
	}
	return false, nil
}

func (u *User) ExistsEmail(ctx context.Context, user *model.User) (bool, error) {
	var data []model.User
	q := datastore.NewQuery(u.kind).Namespace(config.NameSpace).Filter("Email =", user.Email).Limit(1)
	_, err := u.client.GetAll(ctx, q, &data)
	if err != nil {
		if err == datastore.ErrNoSuchEntity {
			return false, nil
		}
		return false, err
	} else if 0 < len(data) {
		log.Println("Email is exists")
		return true, nil
	}
	return false, nil
}
