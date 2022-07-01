package api

import (
	"fhz/api/handlers"
	"github.com/gin-gonic/gin"
)

type Router struct {
	article handlers.ArticleHandel
}

func (r Router) With(engine *gin.Engine) {
	engine.POST("/articles", r.article.FetchArticle)
}

func NewRouter(article handlers.ArticleHandel) *Router {
	return &Router{
		article: article,
	}
}
