package model

type CategoryModel struct {
	CateId uint `json:"cate_id"`
	Name string `json:"name"`
	Sort string `json:"sort"`
	Articles []*ArticleModel `json:"articles" gorm:"foreignKey:CateId;references:CateId"`
}

// TableName 自定义表名
func (CategoryModel) TableName() string {
	return "blog_category"
}

