package article_api

import (
	"blog/global"
	"blog/model"
	"blog/utils/response"
	"fmt"
	"github.com/gin-gonic/gin"
)

type ArticleParams struct {
	ArticleId     int `form:"article_id"`
}

func (ArticleApi) Article(c *gin.Context) {
	var params ArticleParams
	err := c.ShouldBind(&params)
	if err != nil {
		fmt.Println(err)
		response.ErrorWithMessage(err.Error(), c)
		return
	}
	var articleModel model.ArticleModel
	// count := global.DB.Scopes(db.Paginate(ArticleParams.Page, ArticleParams.PageSize)).Find(&articleModel).RowsAffected
	global.DB.Where("article_id = ?", params.ArticleId).First(&articleModel)

	response.SuccessWithData(articleModel, c)
}
