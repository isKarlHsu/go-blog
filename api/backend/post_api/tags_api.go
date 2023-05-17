package post_api

import (
	"blog/global"
	"blog/model"
	"blog/utils/response"
	"github.com/gin-gonic/gin"
)

func (PostApi) Tags(c *gin.Context)  {
	var TagModel []model.TagModel
	global.DB.Find(&TagModel)
	response.SuccessWithData(TagModel, c)
}