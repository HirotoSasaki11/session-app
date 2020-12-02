package main

import (
	"log"
	"net/http"
	"session-sample/server/adapter/driver/datastore"
	"session-sample/server/adapter/driver/redis"
	"session-sample/server/adapter/router"
	"session-sample/server/application"
	"session-sample/server/registry"
)

func main() {
	log.Println("start application")
	resources := registry.NewResouceForRedisApplication()
	defer resources.FinilizeForRedisApplication()

	p := &router.Provide{
		User: &router.User{
			User: &application.User{
				User: datastore.ProvideUser(resources.DsClient),
			},
		},
		Session: &router.Session{
			Session: &application.Session{
				Session: redis.ProveideSession(resources.RedisClient),
				User:    datastore.ProvideUser(resources.DsClient),
			},
		},
	}
	err := http.ListenAndServe(":8080", router.NewRouter(resources, *p))
	if err != nil {
		log.Fatal("ListenAndServe", err)
	}
	log.Println("start application")
}
