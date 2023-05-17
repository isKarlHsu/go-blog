package model

type TagModel struct {
	TagId uint `json:"tag_id" gorm:"primarykey"`
	Name string `json:"name"`
}

// TableName 自定义表名
func (TagModel) TableName() string {
	return "blog_tag"
}