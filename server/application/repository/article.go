package repository

import (
	"session-sample/server/application/model"
)

type Article interface {
	Get(id string) (*model.Article, error)
	Create(model.Article) error
}
