package article_api

import (
	"blog/global"
	"blog/model"
	"blog/utils/db"
	"blog/utils/response"
	"fmt"
	"github.com/gin-gonic/gin"
	"math"
)

type ArticleListParams struct {
	Page     int `json:"page"`
	PageSize int `json:"page_size"`
}

func (ArticleApi) Articles(c *gin.Context) {
	var params ArticleListParams
	err := c.ShouldBind(&params)
	if err != nil {
		fmt.Println(err)
		response.ErrorWithMessage(err.Error(), c)
		return
	}
	var articleModel []model.ArticleModel
	// count := global.DB.Scopes(db.Paginate(ArticleParams.Page, ArticleParams.PageSize)).Find(&articleModel).RowsAffected
	global.DB.Order("article_id Desc").Find(&articleModel)
	count := int64(len(articleModel))
	global.DB.Scopes(db.Paginate(params.Page, params.PageSize)).Order("article_id Desc").Find(&articleModel)
	data := db.Page{
		CurrentPage: int64(params.Page),
		PageSize:    int64(params.PageSize),
		Total: 		 count,
		Pages:       int64(math.Ceil(float64(count) / float64(params.PageSize))),
		List:        articleModel,
	}
	// fmt.Println(articleModel)
	response.SuccessWithData(data, c)
}
