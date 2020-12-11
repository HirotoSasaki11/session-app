package redis

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"session-sample/server/application/model"
	"session-sample/server/config"
	"session-sample/server/lib"

	"github.com/go-redis/redis/v8"
	"github.com/rbcervilla/redisstore/v8"
)

type Session struct {
	client *redis.Client
}

func ProveideSession(client *redis.Client) *Session {
	return &Session{client: client}
}

func (s *Session) Get(ctx context.Context, w http.ResponseWriter, r *http.Request) (*model.User, error) {
	user := new(model.User)
	store, err := redisstore.NewRedisStore(ctx, s.client)
	if err != nil {
		log.Fatal("failed to create redis store: ", err)
	}

	// Get a session.
	sess, err := store.Get(r, config.SessionKey)

	if err != nil {
		log.Fatal("failed to create redis store: ", err)
		return nil, err
	}

	_, ok := sess.Values[sess.ID]
	if !ok {
		log.Println("not found session id.")
		return nil, nil
	}
	err = json.Unmarshal([]byte(sess.Values[sess.ID].(string)), user)
	if err != nil {
		return nil, err
	}
	log.Println(user)
	return user, nil
}

func (s *Session) Set(ctx context.Context, w http.ResponseWriter, r *http.Request, user *model.User) error {
	store, err := redisstore.NewRedisStore(ctx, s.client)
	if err != nil {
		log.Fatal("failed to create redis store: ", err)
		return err
	}

	sess, err := store.New(r, config.SessionKey)
	if err != nil {
		return err
	}
	sess.ID = lib.NewStringID()

	data, err := json.Marshal(user)
	if err != nil {
		return err
	}
	sess.Values[sess.ID] = string(data)
	if err := sess.Save(r, w); err != nil {
		log.Printf("Error saving session: %v", err)
		return err
	}
	return nil
}
