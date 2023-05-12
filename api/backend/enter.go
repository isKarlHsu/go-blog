package backend

import (
	"blog/api/backend/post_api"
	"blog/api/backend/upload_api"
	"blog/api/backend/user_api"
)

type BackendApiGroup struct {
	UserApi    user_api.UserApi
	PostApi post_api.PostApi
	UploadApi upload_api.UploadApi
}

var BackendApi = new(BackendApiGroup)
