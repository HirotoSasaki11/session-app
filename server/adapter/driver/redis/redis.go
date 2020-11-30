package redis

import (
	"fmt"
	"session-sample/server/config"

	"github.com/go-redis/redis/v8"
)

func Connection() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%s", config.RedisHost, config.RedisPort),
	})

	// con, err := redis.Dial("tcp", fmt.Sprintf("%s:%s", config.RedisHost, config.RedisPort))
	// if err != nil {
	// 	panic(err)
	// }
	// return con
	return client
}
