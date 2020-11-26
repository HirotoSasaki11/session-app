package memcache

import (
	"fmt"
	"os"

	"github.com/bradfitz/gomemcache/memcache"
)

func Connection() *memcache.Client {
	host := os.Getenv("MEMCACHE_HOST")
	port := os.Getenv("MEMCACHE_PORT")
	mc := memcache.New(fmt.Sprintf("%s:%s", host, port))
	return mc
}
