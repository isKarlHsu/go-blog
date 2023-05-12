package post_api

import (
	"blog/global"
	"blog/model"
	"blog/utils/response"
	"github.com/gin-gonic/gin"
)

func (PostApi) Cates(c *gin.Context) {
	var cateModel []model.CategoryModel
	global.DB.Preload("Articles").Find(&cateModel)
	response.SuccessWithData(cateModel, c)
}
