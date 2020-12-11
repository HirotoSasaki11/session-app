package application

import (
	"context"
	"session-sample/server/application/model"
	"session-sample/server/application/repository"
)

type Article struct {
	Article repository.Article
}

func (a *Article) GetByID(ctx context.Context, id string) (*model.Article, error) {
	return a.Article.GetByID(ctx, id)
}

func (a *Article) Create(ctx context.Context, article *model.Article) (*model.Article, error) {
	return a.Article.Create(ctx, article)
}
