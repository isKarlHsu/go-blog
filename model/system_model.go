package model

type SystemModel struct {
	SysId uint `json:"sys_id" gorm:"primarykey"`
	Name string `json:"name"`
	Value string `json:"value"`
}

// TableName 自定义表名
func (SystemModel) TableName() string {
	return "blog_system"
}