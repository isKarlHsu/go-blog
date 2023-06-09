package user_api

import (
	"blog/global"
	"blog/model"
	"blog/utils/response"
	"fmt"
	"github.com/gin-gonic/gin"
)

func (UserApi) About(c *gin.Context)  {
	var systemModel model.SystemModel
	global.DB.Where("name = ?", "about").First(&systemModel)
	response.SuccessWithData(systemModel, c)
}

type AboutEditParams struct {
	Content     string `json:"content"`
}
func (UserApi) AboutEdit(c *gin.Context)  {
	var params AboutEditParams
	err := c.ShouldBind(&params)
	if err != nil {
		fmt.Println(err)
		response.ErrorWithMessage(err.Error(), c)
		return
	}
	var systemModel model.SystemModel
	global.DB.Where("name = ?", "about").First(&systemModel)
	systemModel.Value = params.Content
	global.DB.Save(&systemModel)
	response.SuccessWithMessage("操作成功", c)
}
