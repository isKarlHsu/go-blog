package model

type ArticleTagModel struct {
	ArticleTagId uint   `gorm:"primarykey" json:"article_tag_id"`
	ArticleId int `json:"article_id"`
	TagId int `json:"tag_id"`
	Tag *TagModel `json:"tag" gorm:"foreignKey:TagId;references:TagId"`
}

// TableName 自定义表名
func (ArticleTagModel) TableName() string {
	return "blog_article_tag"
}