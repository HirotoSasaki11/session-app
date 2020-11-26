package redis

import (
	"fmt"
	"os"

	"github.com/gomodule/redigo/redis"
)

func Connection() redis.Conn {
	host := os.Getenv("RADIS_HOST")
	port := os.Getenv("RADIS_PORT")
	con, err := redis.Dial("tcp", fmt.Sprintf("%s:%s", host, port))
	if err != nil {
		panic(err)
	}
	return con
}
