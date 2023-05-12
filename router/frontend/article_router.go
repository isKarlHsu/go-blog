package frontend

import (
	"blog/api/frontend"
	"github.com/gin-gonic/gin"
)

func ArticleRouter(router *gin.RouterGroup) {
	ArticleApi := frontend.FrontendApi.ArticleApi
	router.POST("/articleList", ArticleApi.ArticleList)
	router.POST("/article", ArticleApi.Article)
}
