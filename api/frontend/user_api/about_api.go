package user_api

import (
	"blog/global"
	"blog/model"
	"blog/utils/response"
	"github.com/gin-gonic/gin"
)

func (UserApi) About(c *gin.Context)  {
	var systemModel model.SystemModel
	global.DB.Where("name = ?", "about").First(&systemModel)
	response.SuccessWithData(systemModel, c)
}
