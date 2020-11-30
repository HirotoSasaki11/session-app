package config

import "os"

var (
	ProjectID    = os.Getenv("PROJECTID")
	NameSpace    = os.Getenv("NAMEAPACE")
	MemCacheHost = os.Getenv("MEMCACHE_HOST")
	MemCachePort = os.Getenv("MEMCACHE_PORT")
	RedisHost    = os.Getenv("RADIS_HOST")
	RedisPort    = os.Getenv("RADIS_PORT")
	Credentials  = os.Getenv("CREDENTIALS")
	SessionKey   = os.Getenv("SESSIONKEY")
)
