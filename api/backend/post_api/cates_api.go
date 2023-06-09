package post_api

import (
	"blog/global"
	"blog/model"
	"blog/utils/response"
	"fmt"
	"github.com/gin-gonic/gin"
)

func (PostApi) Cates(c *gin.Context) {
	var cateModel []model.CategoryModel
	global.DB.Find(&cateModel)
	response.SuccessWithData(cateModel, c)
}

type CateEditParams struct {
	CateId int8 `json:"cate_id"`
	Name string `json:"name"`
}

func (PostApi) CateEdit(c *gin.Context) {
	var params CateEditParams
	err := c.ShouldBind(&params)
	if err != nil {
		fmt.Println(err)
		response.ErrorWithMessage(err.Error(), c)
		return
	}
	var cateModel model.CategoryModel
	global.DB.Where("cate_id = ?", params.CateId).First(&cateModel)
	cateModel.Name = params.Name
	global.DB.Where("cate_id = ?", params.CateId).Save(&cateModel)
	response.SuccessWithMessage("编辑成功", c)

}
