package registry

import (
	ds "session-sample/server/adapter/driver/datastore"
	r "session-sample/server/adapter/driver/redis"

	"cloud.google.com/go/datastore"
	"github.com/gomodule/redigo/redis"
)

type Resource struct {
	ConectRedis redis.Conn
	DsClient    *datastore.Client
}

func NewResouce() Resource {
	return Resource{
		ConectRedis: r.Connection(),
		DsClient:    ds.NewClient(),
	}
}

func (r *Resource) Finilize() {
	r.ConectRedis.Close()
	r.DsClient.Close()
}
