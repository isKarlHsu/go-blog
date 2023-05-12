package command

import (
	"blog/global"
	"blog/model"
)

func Migrations() {
	err := global.DB.Migrator().AutoMigrate(
		&model.UserModel{},
		&model.ArticleModel{},
		&model.CategoryModel{},
	)
	if err != nil {
		global.Logger.Error("生成数据表失败")
		return
	}
	global.Logger.Info("生成数据表成功")
}
