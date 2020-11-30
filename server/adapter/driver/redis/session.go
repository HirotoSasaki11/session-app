package redis

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"session-sample/server/application/model"
	"session-sample/server/config"

	"github.com/go-redis/redis/v8"
	"github.com/rbcervilla/redisstore/v8"
)

type Session struct {
	client *redis.Client
}

// func ProveideSession(con redis.Conn) *Session {
// 	return &Session{con: con}
// }
func ProveideSession(client *redis.Client) *Session {
	return &Session{client: client}
}

// func (s Session) Get(key string) (*model.User, error) {
// var user *model.User
// value, err := s.store.Get()
// if err != nil {
// 	return nil, err
// }

// err = json.Unmarshal(value.([]byte), &user)
// if err != nil {
// 	return nil, err
// }
// return user, nil
//}

// func (s Session) Get(key string) (*model.User, error) {
// 	var user *model.User
// 	value, err := s.con.Do("Get", key)
// 	if err != nil {
// 		return nil, err
// 	}

// 	err = json.Unmarshal(value.([]byte), &user)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return user, nil
// }
// func (s Session) Set(key, value string) error {
// 	_, err := s.con.Do("Set", key, value)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

func (s Session) Get(ctx context.Context, w http.ResponseWriter, r *http.Request, id string) (*model.User, error) {
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
	if sess.Values[id] == nil {
		log.Println("not found session id.")
		return nil, nil
	}
	err = json.Unmarshal([]byte(sess.Values[id].(string)), user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s Session) Set(w http.ResponseWriter, r *http.Request, user *model.User) error {
	store, err := redisstore.NewRedisStore(context.Background(), s.client)
	if err != nil {
		log.Fatal("failed to create redis store: ", err)
		return err
	}
	sess, err := store.New(r, config.SessionKey)
	if err != nil {
		return err
	}
	data, err := json.Marshal(user)
	if err != nil {
		return err
	}
	sess.Values[user.ID] = string(data)
	if err := sess.Save(r, w); err != nil {
		log.Printf("Error saving session: %v", err)
		return err
	}
	return nil
}
