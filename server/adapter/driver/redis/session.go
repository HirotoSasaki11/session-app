package redis

import (
	"encoding/json"
	"session-sample/server/application/model"

	"github.com/gomodule/redigo/redis"
)

type Session struct {
	con redis.Conn
}

func ProveideSession(con redis.Conn) *Session {
	return &Session{con: con}
}
func (s Session) Get(key string) (*model.User, error) {
	var user *model.User
	value, err := s.con.Do("Get", key)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(value.([]byte), &user)
	if err != nil {
		return nil, err
	}
	return user, nil
}
func (s Session) Set(key, value string) error {
	_, err := s.con.Do("Set", key, value)
	if err != nil {
		return err
	}
	return nil
}
