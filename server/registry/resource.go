package registry

import (
	ds "session-sample/server/adapter/driver/datastore"
	"session-sample/server/adapter/driver/memcache"
	r "session-sample/server/adapter/driver/redis"

	"cloud.google.com/go/datastore"
	"github.com/go-redis/redis/v8"
	"github.com/memcachier/mc"
)

type Resource struct {
	Memcached   *mc.Client
	RedisClient *redis.Client
	DsClient    *datastore.Client
}

func NewResouceForRedisApplication() *Resource {
	return &Resource{
		RedisClient: r.Connection(),
		DsClient:    ds.NewClient(),
	}
}

func (r *Resource) FinilizeForRedisApplication() {
	r.RedisClient.Close()
	r.DsClient.Close()
}

func NewResouceForMemCachedApplication() *Resource {
	return &Resource{
		Memcached: memcache.Connection(),
		DsClient:  ds.NewClient(),
	}
}

func (r *Resource) FinilizeForMemCachedApplication() {
	r.DsClient.Close()
}
