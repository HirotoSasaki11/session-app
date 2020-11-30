package memcache

import (
	"fmt"
	"session-sample/server/config"

	"github.com/bradfitz/gomemcache/memcache"
)

func Connection() *memcache.Client {
	mc := memcache.New(fmt.Sprintf("%s:%s", config.MemCacheHost, config.MemCachePort))
	return mc
}
