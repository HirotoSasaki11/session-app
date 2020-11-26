package main

import (
	"log"
	"net/http"
	"session-sample/server/adapter/driver/redis"
	"session-sample/server/adapter/router"
	"session-sample/server/application"
	"session-sample/server/registry"
)

func main() {
	resources := registry.NewResouce()
	p := &router.Provide{
		// User: ,
		Session: &router.Session{
			Session: &application.Session{
				Session: redis.ProveideSession(resources.ConectRedis),
			},
		},
	}
	err := http.ListenAndServe(":8080", router.NewRouter(resources, *p))
	if err != nil {
		log.Fatal("ListenAndServe", err)
	}
}
