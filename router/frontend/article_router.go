package frontend

import (
	"blog/api/frontend"
	"github.com/gin-gonic/gin"
)

func ArticleRouter(router *gin.RouterGroup) {
	ArticleApi := frontend.FrontendApi.ArticleApi
	router.POST("/articles", ArticleApi.Articles)
	router.POST("/article", ArticleApi.Article)
}
