package frontend

import (
	"blog/api/frontend"
	"github.com/gin-gonic/gin"
)

func PostRouter(router *gin.RouterGroup) {
	PostApi := frontend.FrontendApi.PostApi
	router.POST("/posts", PostApi.Post)
	router.POST("/cates", PostApi.Cates)
}
