package backend

import (
	"blog/api/backend"
	"github.com/gin-gonic/gin"
)

func PostRouter(router *gin.RouterGroup) {
	PostApi := backend.BackendApi.PostApi
	router.POST("/cates", PostApi.Cates)
	router.POST("/tags", PostApi.Tags)
	router.POST("/articles", PostApi.Articles)
	router.POST("/article", PostApi.Article)
	router.POST("/articleEdit", PostApi.ArticleEdit)
	router.POST("/articleAdd", PostApi.ArticleAdd)
}
