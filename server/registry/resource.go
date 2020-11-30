package registry

import (
	ds "session-sample/server/adapter/driver/datastore"
	r "session-sample/server/adapter/driver/redis"

	"cloud.google.com/go/datastore"
	"github.com/go-redis/redis/v8"
)

type Resource struct {
	RedisClient *redis.Client
	DsClient    *datastore.Client
}

func NewResouceForRedisApplication() Resource {
	return Resource{
		RedisClient: r.Connection(),
		DsClient:    ds.NewClient(),
	}
}

func (r *Resource) FinilizeForRedisApplication() {
	r.ConectRedis.Close()
	r.DsClient.Close()
}
