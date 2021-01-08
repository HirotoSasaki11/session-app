package memcache

import (
	"session-sample/server/config"

	"github.com/memcachier/mc"
)

func Connection() *mc.Client {
	mc := mc.NewMC(config.MemCacheHost, config.MemCachePort, "")
	return mc
}
