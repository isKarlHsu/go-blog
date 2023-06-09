package frontend

import (
	"blog/api/frontend"
	"github.com/gin-gonic/gin"
)

func UserRouter(router *gin.RouterGroup) {
	UserApiApp := frontend.FrontendApi.UserApi
	router.POST("/about", UserApiApp.About)
}
