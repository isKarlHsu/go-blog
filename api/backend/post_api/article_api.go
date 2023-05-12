package post_api

import (
	"blog/global"
	"blog/model"
	"blog/utils/db"
	"blog/utils/response"
	"fmt"
	"github.com/gin-gonic/gin"
	"math"
)

type ArticlesParams struct {
	Page     int `json:"page"`
	PageSize int `json:"page_size"`
	Title string `json:"title"`
	CateId int `json:"cate_id"`
}

func (PostApi) Articles(c *gin.Context) {
	var params ArticlesParams
	err := c.ShouldBind(&params)
	if err != nil {
		fmt.Println(err)
		response.ErrorWithMessage(err.Error(), c)
		return
	}
	var articleModel []model.ArticleModel
	global.DB.Order("article_id Desc").Find(&articleModel)
	count := int64(len(articleModel))
	global.DB.
		Scopes(model.ArticleFilter(params)).
		Scopes(db.Paginate(params.Page, params.PageSize)).
		Preload("Category").
		Order("article_id Desc").
		Find(&articleModel)
	data := db.Page{
		CurrentPage: int64(params.Page),
		PageSize:    int64(params.PageSize),
		Total: 		 count,
		Pages:       int64(math.Ceil(float64(count) / float64(params.PageSize))),
		List:        articleModel,
	}
	response.SuccessWithData(data, c)
}

type ArticleParams struct {
	ArticleId int `json:"article_id""`
}

func (PostApi) Article(c *gin.Context) {
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

