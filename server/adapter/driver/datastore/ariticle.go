package datastore

import (
	"context"
	"session-sample/server/application/model"
	"session-sample/server/lib"

	"cloud.google.com/go/datastore"
)

type Ariticle struct {
	client *datastore.Client
	kind   string
}

func ProvideArticle(client *datastore.Client) *Ariticle {
	return &Ariticle{client: client, kind: model.ArticleEntityName}
}

func (a *Ariticle) Create(ctx context.Context, article *model.Article) (*model.Article, error) {
	article.ID = lib.NewStringID()
	key := NewKey(a.kind, article.ID)

	_, err := a.client.Put(ctx, key, article)
	if err != nil {
		return nil, err
	}
	return article, nil
}

func (a *Ariticle) GetByID(ctx context.Context, id string) (*model.Article, error) {
	article := new(model.Article)
	err := a.client.Get(ctx, NewKey(a.kind, id), article)
	if err != nil {
		return nil, err
	}
	return article, nil
}
