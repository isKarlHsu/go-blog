package backend

import (
	"blog/api/backend"
	"github.com/gin-gonic/gin"
)

func PostRouter(router *gin.RouterGroup) {
	PostApi := backend.BackendApi.PostApi
	router.POST("/cates", PostApi.Cates)
	router.POST("/articles", PostApi.Articles)
	router.POST("/article", PostApi.Article)
}
