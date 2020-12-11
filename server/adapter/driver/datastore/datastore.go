package datastore

import (
	"context"
	"session-sample/server/config"

	"cloud.google.com/go/datastore"
	"google.golang.org/api/option"
)

func NewClient() *datastore.Client {
	ctx := context.Background()
	opt := option.WithCredentialsJSON([]byte(config.Credentials))
	client, err := datastore.NewClient(ctx, config.ProjectID, opt)
	if err != nil {
		panic(err)
	}
	return client
}

func NewKey(kind, name string) *datastore.Key {
	return &datastore.Key{
		Name:      name,
		Kind:      kind,
		Namespace: config.NameSpace,
	}
}
