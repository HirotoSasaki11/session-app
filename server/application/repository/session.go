package repository

import (
	"session-sample/server/application/model"
)

type Session interface {
	Get(key string) (*model.User, error)
	Set(key, value string) error
}
