package datastore

import (
	"context"
	"os"
	"session-sample/server/application/model"

	"cloud.google.com/go/datastore"
)

type User struct {
	client *datastore.Client
	kind   string
}

func ProvideUser(client *datastore.Client) *User {
	return &User{client: client}
}

func (u *User) NewKey(kind, name string) *datastore.Key {
	namaSpace := os.Getenv("NAMEAPACE")
	return &datastore.Key{
		Name:      name,
		Kind:      kind,
		Namespace: namaSpace,
	}
}
func (u *User) Create(ctx context.Context, user *model.User) error {
	_, err := u.client.Put(ctx, u.NewKey(u.kind, user.ID), user)
	if err != nil {
		return err
	}
	return nil
}

func (u *User) GetByID(ctx context.Context, user *model.User) (*model.User, error) {
	var data *model.User
	err := u.client.Get(ctx, u.NewKey(u.kind, user.ID), *data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (u *User) GetByName(ctx context.Context, user *model.User) (*model.User, error) {
	var data []model.User
	q := datastore.NewQuery(u.kind).Namespace(os.Getenv("NAMEAPACE")).Filter("Email", user.Name).Limit(1)
	_, err := u.client.GetAll(ctx, q, &data)
	if err != nil {
		return nil, err
	}
	return &data[0], nil
}
