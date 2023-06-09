package frontend

import (
	"blog/api/frontend/article_api"
	"blog/api/frontend/post_api"
	"blog/api/frontend/user_api"
)

type FrontendApiGroup struct {
	ArticleApi article_api.ArticleApi
	PostApi post_api.PostApi
	UserApi user_api.UserApi
}

var FrontendApi = new(FrontendApiGroup)