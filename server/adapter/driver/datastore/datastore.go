package datastore

import (
	"context"
	"os"

	"cloud.google.com/go/datastore"
)

func NewClient() *datastore.Client {
	ctx := context.Background()
	id := os.Getenv("PROJECTID")
	client, err := datastore.NewClient(ctx, id)
	if err != nil {
		panic(err)
	}
	return client
}
