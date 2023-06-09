package router

import (
	"blog/global"
	"blog/router/backend"
	"blog/router/frontend"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	gin.SetMode(global.Config.System.Env)
	router := gin.Default()

	// 访问静态资源
	router.Static("/resource", "resource")

	// 后台接口组
	backendRouterGroup := router.Group("admin")
	backend.UserRouter(backendRouterGroup)
	backend.PostRouter(backendRouterGroup)
	backend.UploadRouter(backendRouterGroup)

	// 前台接口组
	frontendRouterGroup := router.Group("api")
	frontend.UserRouter(frontendRouterGroup)
	frontend.ArticleRouter(frontendRouterGroup)
	frontend.PostRouter(frontendRouterGroup)
	return router
}
