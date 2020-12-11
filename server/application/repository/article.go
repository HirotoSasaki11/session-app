package repository

import (
	"context"
	"session-sample/server/application/model"
)

type Article interface {
	GetByID(ctx context.Context, id string) (*model.Article, error)
	Create(ctx context.Context, article *model.Article) (*model.Article, error)
}
