package router

import (
	"context"
	"log"
	"net/http"
	"session-sample/server/application"
	"session-sample/server/application/model"
	"session-sample/server/lib"

	"github.com/go-chi/chi"
)

type Article struct {
	Article *application.Article
}

func (a *Article) GetByID(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	id := chi.URLParam(r, "article_id")
	article, err := a.Article.GetByID(ctx, id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	} else if article == nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("not found article"))
		return
	}
	lib.Json(w, article)
	w.WriteHeader(http.StatusOK)
	return
}

func (a *Article) Create(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	article := new(model.Article)
	err := lib.BodyToJson(r, article)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	log.Println(article)
	article, err = a.Article.Create(ctx, article)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = lib.Json(w, article)
	if err != nil {
		return
	}
	w.WriteHeader(http.StatusCreated)
}
