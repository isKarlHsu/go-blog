package backend

import (
	"blog/api/backend"
	"github.com/gin-gonic/gin"
)

func UploadRouter(router *gin.RouterGroup) {
	UploadApi := backend.BackendApi.UploadApi
	router.POST("/upload/image", UploadApi.Image)
}