package memcache

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"session-sample/server/application/model"
	"session-sample/server/config"
	"session-sample/server/lib"

	gsm "github.com/bradleypeabody/gorilla-sessions-memcache"
	"github.com/memcachier/mc"
)

type Session struct {
	client *mc.Client
}

func ProveideSession(client *mc.Client) *Session {
	return &Session{client: client}
}
func (s *Session) Get(ctx context.Context, w http.ResponseWriter, r *http.Request) (*model.User, error) {
	var user *model.User

	store := gsm.NewMemcacherStore(s.client, "", []byte(config.SessionKey))
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
	store := gsm.NewMemcacherStore(s.client, "", []byte(config.SessionKey))
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
