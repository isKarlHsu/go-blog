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
		Preload("ArticleTag").
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
	global.DB.Where("article_id = ?", params.ArticleId).
		Preload("Category").
		Preload("ArticleTag.Tag").
		First(&articleModel)

	response.SuccessWithData(articleModel, c)
}

type ArticleAddParams struct {
	Title string `json:"title"`
	CateId int8 `json:"cate_id"`
	Type int8 `json:"type"`
	Author string `json:"author"`
	Sources string `json:"sources"`
	Abstract string `json:"abstract"`
	Content string `json:"content"`
	Tags []Tags `json:"tags"`
}

type Tags struct {
	TagId int `json:"tag_id"`
	Name string `json:"name"`
}
func (PostApi) ArticleAdd(c *gin.Context) {
	var params ArticleAddParams
	err := c.ShouldBind(&params)
	if err != nil {
		fmt.Println(err)
		response.ErrorWithMessage(err.Error(), c)
		return
	}

	var article model.ArticleModel
	article.Title = params.Title
	article.CateId = params.CateId
	article.Type = params.Type
	article.Author = params.Author
	article.Sources = params.Sources
	article.Abstract = params.Abstract
	article.Content = params.Content
	global.DB.Create(&article)

	global.Logger.Println(params.Tags)

	var articleTag []model.ArticleTagModel

	for _, tag := range params.Tags {
		// 新增的 tag
		if tag.TagId == 0 {
			tagModel := model.TagModel{
				Name:  tag.Name,
			}
			global.DB.Create(&tagModel)
			tag.TagId = int(tagModel.TagId)
		}

		articleTag = append(articleTag, model.ArticleTagModel{
			ArticleId: int(article.ArticleId),
			TagId:     tag.TagId,
		})
	}
	global.DB.Create(&articleTag)
	// response.SuccessWithData(article, c)
	response.SuccessWithMessage("发布成功", c)
}


type ArticleEditParams struct {
	ArticleId int `json:"article_id"`
	Title string `json:"title"`
	CateId int8 `json:"cate_id"`
	Type int8 `json:"type"`
	Author string `json:"author"`
	Sources string `json:"sources"`
	Abstract string `json:"abstract"`
	Content string `json:"content"`
	Tags []Tags `json:"tags"`
}
func (PostApi) ArticleEdit(c *gin.Context) {
	var params ArticleEditParams
	err := c.ShouldBind(&params)
	if err != nil {
		fmt.Println(err)
		response.ErrorWithMessage(err.Error(), c)
		return
	}
	var article model.ArticleModel
	global.DB.Where("article_id = ?", params.ArticleId).First(&article)
	article.Title = params.Title
	article.CateId = params.CateId
	article.Type = params.Type
	article.Author = params.Author
	article.Sources = params.Sources
	article.Abstract = params.Abstract
	article.Content = params.Content
	global.DB.Save(&article)

	global.Logger.Println(params.Tags)

	var articleTagModel model.ArticleTagModel
	var articleTag []model.ArticleTagModel

	for _, tag := range params.Tags {
		// 新增的 tag
		if tag.TagId == 0 {
			tagModel := model.TagModel{
				Name:  tag.Name,
			}
			global.DB.Create(&tagModel)
			tag.TagId = int(tagModel.TagId)
		}

		articleTag = append(articleTag, model.ArticleTagModel{
			ArticleId: int(article.ArticleId),
			TagId:     tag.TagId,
		})
	}
	global.DB.Where("article_id = ?", article.ArticleId).Delete(&articleTagModel)
	global.DB.Create(&articleTag)

	// response.SuccessWithData(article, c)
	response.SuccessWithMessage("编辑成功", c)
}