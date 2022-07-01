package handlers

import (
	"fhz/model"
	"fhz/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ArticleHandel struct {
	service service.IArticleService
}

func NewArticleHandel(articleService service.IArticleService) ArticleHandel {
	return ArticleHandel{service: articleService}
}

func (a *ArticleHandel) FetchArticle(c *gin.Context) {
	req := model.ArticleVo{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": fmt.Sprintf("params error: %v", err),
		})
		return
	}
	list, err := a.service.Fetch(c, req.CreateData, req.Num)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": fmt.Sprintf("Internal server error: %v", err),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    1,
		"message": "success",
		"data":    list,
	})
}
