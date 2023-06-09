package backend

import (
	"blog/api/backend"
	"github.com/gin-gonic/gin"
)

func UserRouter(router *gin.RouterGroup) {
	UserApiApp := backend.BackendApi.UserApi
	router.GET("/login", UserApiApp.UserLogin)
	router.POST("/about", UserApiApp.About)
	router.POST("/aboutEdit", UserApiApp.AboutEdit)
}
