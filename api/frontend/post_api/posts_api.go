package post_api

import (
	"blog/global"
	"blog/model"
	"blog/utils/response"
	"github.com/gin-gonic/gin"
)

type NewArticleModel struct {
	model.ArticleModel
	Datetime string `json:"datetime"`
}

func (PostApi) Post(c *gin.Context){
	var articleModel []NewArticleModel
	global.DB.Select("article_id", "title", "DATE_FORMAT(created_at, '%Y') as datetime").Order("created_at DESC").Find(&articleModel)

	data := make(map[string][]NewArticleModel)

	for _, article := range articleModel {
		data[article.Datetime] = append(data[article.Datetime], article)
	}

	response.SuccessWithData(data, c)
}
