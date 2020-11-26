package memcache

import (
	"encoding/json"
	"session-sample/server/application/model"

	"github.com/bradfitz/gomemcache/memcache"
)

type Session struct {
	mc *memcache.Client
}

func (s Session) Get(key string) (*model.User, error) {
	var user *model.User

	item, err := s.mc.Get(key)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(item.Value, user)
	if err != nil {
		return nil, err
	}
	return user, nil
}
func (s Session) Set(key, value string) error {
	item := &memcache.Item{
		Key:   key,
		Value: []byte(value),
	}
	err := s.mc.Set(item)
	if err != nil {
		return err
	}
	return nil
}
