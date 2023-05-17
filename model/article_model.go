package model

import (
	"gorm.io/gorm"
	"reflect"
)

type ArticleModel struct {
	ArticleId uint   `gorm:"primarykey" json:"article_id"`
	Title     string `gorm:"size:255" json:"title"`
	CateId    int8   `gorm:"size:255" json:"cate_id"`
	Type      int8   `gorm:"size:20" json:"type"`
	Sort      int8   `gorm:"size:20" json:"sort"`
	Abstract  string `gorm:"size:2048" json:"abstract"`
	Author    string `gorm:"size:64" json:"author"`
	Sources   string `gorm:"size:1024" json:"sources"`
	Content   string `gorm:"type:longtext" json:"content"`
	Category *CategoryModel `json:"category" gorm:"foreignKey:CateId;references:CateId"`
	ArticleTag *[]ArticleTagModel `json:"article_tag" gorm:"foreignKey:ArticleId;references:ArticleId"`
	Timestamp
}

// TableName 自定义表名
func (ArticleModel) TableName() string {
	return "blog_article"
}

func ArticleFilter(params any) func (db *gorm.DB) *gorm.DB {
	return func (db *gorm.DB) *gorm.DB {
		immutableV := reflect.ValueOf(params)
		title := immutableV.FieldByName("Title").String()
		if title != "" {
			db.Where("title like ?", "%"+title+"%")
		}
		cateId := immutableV.FieldByName("CateId").Int()
		if cateId != 0 {
			db.Where("cate_id =  ?", cateId)
		}
		return db
	}
}