package frontend

import (
	"blog/api/frontend/article_api"
	"blog/api/frontend/post_api"
)

type FrontendApiGroup struct {
	ArticleApi article_api.ArticleApi
	PostApi post_api.PostApi
}

var FrontendApi = new(FrontendApiGroup)
